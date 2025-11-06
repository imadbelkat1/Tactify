package tests

import (
	"context"
	"reflect"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	eventService "github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestEventsService(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	service := &eventService.EventsService{
		Config:   config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
	}

	var laLigaSeasonIDs []int
	var premierLeagueSeasonIDs []int
	var leagueIDs []int

	// Extract La Liga season IDs
	ligaSeason := reflect.ValueOf(service.Config.SofascoreApi.LaLigaSeasonsIDs)
	for i := 0; i < ligaSeason.NumField(); i++ {
		laLigaSeasonIDs = append(laLigaSeasonIDs, int(ligaSeason.Field(i).Int()))
	}

	// Extract Premier League season IDs
	plSeason := reflect.ValueOf(service.Config.SofascoreApi.PremierLeagueSeasonIDs)
	for i := 0; i < plSeason.NumField(); i++ {
		premierLeagueSeasonIDs = append(premierLeagueSeasonIDs, int(plSeason.Field(i).Int()))
	}

	// Extract league IDs
	league := reflect.ValueOf(service.Config.SofascoreApi.LeaguesID)
	for i := 0; i < league.NumField(); i++ {
		leagueIDs = append(leagueIDs, int(league.Field(i).Int()))
	}

	// Rate limiting: max 5 concurrent requests to avoid overwhelming proxy
	const maxConcurrent = 5
	sem := semaphore.NewWeighted(maxConcurrent)

	// Create errgroup for concurrent execution
	g, gCtx := errgroup.WithContext(ctx)

	// Helper function with retry logic
	updateWithRetry := func(ctx context.Context, seasonId, leagueId, round int, leagueName string) error {
		const maxRetries = 3
		const retryDelay = 2 * time.Second

		for attempt := 1; attempt <= maxRetries; attempt++ {
			if err := service.UpdateRoundMatches(ctx, seasonId, leagueId, round); err != nil {
				if attempt < maxRetries {
					t.Logf("Retry %d/%d for %s SeasonID: %d, Round: %d (error: %v)",
						attempt, maxRetries, leagueName, seasonId, round, err)
					time.Sleep(retryDelay)
					continue
				}
				return err
			}
			return nil
		}
		return nil
	}

	// Process each league
	for _, leagueId := range leagueIDs {
		leagueId := leagueId // Capture loop variable

		if leagueId == service.Config.SofascoreApi.LeaguesID.LaLiga {
			// Process La Liga seasons
			for _, seasonId := range laLigaSeasonIDs {
				seasonId := seasonId // Capture loop variable

				for round := 1; round <= 38; round++ {
					round := round // Capture loop variable

					g.Go(func() error {
						// Acquire semaphore to limit concurrency
						if err := sem.Acquire(gCtx, 1); err != nil {
							return err
						}
						defer sem.Release(1)

						t.Logf("Fetching LaLiga SeasonID: %d, LeagueID: %d, Round: %d", seasonId, leagueId, round)

						if err := updateWithRetry(gCtx, seasonId, leagueId, round, "LaLiga"); err != nil {
							t.Errorf("Error updating matches for LaLiga SeasonID %d, Round %d: %v", seasonId, round, err)
							return err
						}
						return nil
					})
				}
			}
		} else if leagueId == service.Config.SofascoreApi.LeaguesID.PremierLeague {
			// Process Premier League seasons
			for _, seasonId := range premierLeagueSeasonIDs {
				seasonId := seasonId // Capture loop variable

				for round := 1; round <= 38; round++ {
					round := round // Capture loop variable

					g.Go(func() error {
						// Acquire semaphore to limit concurrency
						if err := sem.Acquire(gCtx, 1); err != nil {
							return err
						}
						defer sem.Release(1)

						t.Logf("Fetching Premier League SeasonID: %d, LeagueID: %d, Round: %d", seasonId, leagueId, round)

						if err := updateWithRetry(gCtx, seasonId, leagueId, round, "Premier League"); err != nil {
							t.Errorf("Error updating matches for Premier League SeasonID %d, Round %d: %v", seasonId, round, err)
							return err
						}
						return nil
					})
				}
			}
		}
	}

	t.Log("Calling Sofascore API concurrently...")

	// Wait for all goroutines to complete
	if err := g.Wait(); err != nil {
		t.Fatalf("One or more UpdateRoundMatches failed: %v", err)
	}

	t.Log("Sofascore API Publish test completed successfully")
}
