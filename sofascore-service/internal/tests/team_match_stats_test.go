package tests

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	teamMatchStatservice "github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestTeamMatchStatsService(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	service := &teamMatchStatservice.TeamMatchStatsService{
		Event: &teamMatchStatservice.EventsService{
			Config: config.LoadConfig(),
			Client: sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		},
		Config:   *config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
	}

	log.Println("Calling SofaScore API with concurrent processing...")

	// Extract season IDs
	var laLigaSeasonIDs []int
	var premierLeagueSeasonIDs []int

	ligaSeason := reflect.ValueOf(service.Config.SofascoreApi.LaLigaSeasonsIDs)
	for i := 0; i < ligaSeason.NumField(); i++ {
		laLigaSeasonIDs = append(laLigaSeasonIDs, int(ligaSeason.Field(i).Int()))
	}

	plSeason := reflect.ValueOf(service.Config.SofascoreApi.PremierLeagueSeasonIDs)
	for i := 0; i < plSeason.NumField(); i++ {
		premierLeagueSeasonIDs = append(premierLeagueSeasonIDs, int(plSeason.Field(i).Int()))
	}

	// Build work items
	type workItem struct {
		seasonID int
		leagueID int
		round    int
	}

	var work []workItem

	for _, seasonID := range laLigaSeasonIDs {
		for round := 1; round <= 38; round++ {
			work = append(work, workItem{seasonID, service.Config.SofascoreApi.LeaguesID.LaLiga, round})
		}
	}

	for _, seasonID := range premierLeagueSeasonIDs {
		for round := 1; round <= 38; round++ {
			work = append(work, workItem{seasonID, service.Config.SofascoreApi.LeaguesID.PremierLeague, round})
		}
	}

	// Process with worker pool
	start := time.Now()
	const numWorkers = 10 // Tune this based on API rate limits

	workChan := make(chan workItem, len(work))
	errChan := make(chan error, len(work))

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for item := range workChan {
				log.Printf("[Worker %d] Processing season %d, league %d, round %d",
					workerID, item.seasonID, item.leagueID, item.round)

				err := service.UpdateLeagueMatchStats(ctx, item.seasonID, item.leagueID, item.round)
				if err != nil {
					errChan <- fmt.Errorf("season %d, league %d, round %d: %w",
						item.seasonID, item.leagueID, item.round, err)
				}
			}
		}(i)
	}

	// Send work
	for _, item := range work {
		workChan <- item
	}
	close(workChan)

	// Wait for completion
	wg.Wait()
	close(errChan)

	// Collect errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	elapsed := time.Since(start)
	log.Printf("Processing completed in: %v", elapsed)
	log.Printf("Processed %d items with %d workers", len(work), numWorkers)

	if len(errs) > 0 {
		t.Fatalf("Processing failed with %d errors. First error: %v", len(errs), errs[0])
	}

	t.Log("SofaScore API test completed successfully")
}
