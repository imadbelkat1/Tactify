package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	"github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestMatchLineupService(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	eventsService := &services.EventsService{
		Config: config.LoadConfig(),
		Client: sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
	}

	service := &services.MatchLineupService{
		Event:    eventsService,
		Config:   config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
	}

	var laLigaSeasonIDs []int
	var premierLeagueSeasonIDs []int
	var leagueIDs []int

	ligaSeason := reflect.ValueOf(service.Config.SofascoreApi.LaLigaSeasonsIDs)
	for i := 0; i < ligaSeason.NumField(); i++ {
		laLigaSeasonIDs = append(laLigaSeasonIDs, int(ligaSeason.Field(i).Int()))
	}

	plSeason := reflect.ValueOf(service.Config.SofascoreApi.PremierLeagueSeasonIDs)
	for i := 0; i < plSeason.NumField(); i++ {
		premierLeagueSeasonIDs = append(premierLeagueSeasonIDs, int(plSeason.Field(i).Int()))
	}

	league := reflect.ValueOf(service.Config.SofascoreApi.LeaguesID)
	for i := 0; i < league.NumField(); i++ {
		leagueIDs = append(leagueIDs, int(league.Field(i).Int()))
	}

	for _, leagueId := range leagueIDs {
		if leagueId == service.Config.SofascoreApi.LeaguesID.LaLiga {
			for _, seasonId := range laLigaSeasonIDs {
				for i := 1; i <= 38; i++ {
					t.Logf("Fetching LaLiga SeasonID: %d, LeagueID: %d, Round: %d", seasonId, leagueId, i)
					if err := service.UpdatePlayersStats(ctx, seasonId, leagueId, i); err != nil {
						t.Fatalf("Error updating player stats for LaLiga SeasonID %d, Round %d: %v", seasonId, i, err)
					}
				}
			}
		} else if leagueId == service.Config.SofascoreApi.LeaguesID.PremierLeague {
			for _, seasonId := range premierLeagueSeasonIDs {
				for i := 1; i <= 38; i++ {
					t.Logf("Fetching Premier League SeasonID: %d, LeagueID: %d, Round: %d", seasonId, leagueId, i)
					if err := service.UpdatePlayersStats(ctx, seasonId, leagueId, i); err != nil {
						t.Fatalf("Error updating player stats for Premier League SeasonID %d, Round %d: %v", seasonId, i, err)
					}
				}
			}
		}
	}

	t.Log("Calling Sofascore API...")
	t.Log("Sofascore API Publish test completed successfully")
}
