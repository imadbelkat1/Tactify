package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	leagueStandingService "github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestLeagueStandingService(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	service := &leagueStandingService.LeagueStandingService{
		Config:   config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
	}

	var laLigaSeasonIDs []int
	var premierLeagueSeasonIDs []int
	var leagueIDs []int

	ligaS := reflect.ValueOf(service.Config.SofascoreApi.LaLigaSeasonsIDs)
	for i := 0; i < ligaS.NumField(); i++ {
		laLigaSeasonIDs = append(laLigaSeasonIDs, int(ligaS.Field(i).Int()))
	}

	plS := reflect.ValueOf(service.Config.SofascoreApi.PremierLeagueSeasonIDs)
	for i := 0; i < plS.NumField(); i++ {
		premierLeagueSeasonIDs = append(premierLeagueSeasonIDs, int(plS.Field(i).Int()))
	}

	l := reflect.ValueOf(service.Config.SofascoreApi.LeaguesID)
	for i := 0; i < l.NumField(); i++ {
		leagueIDs = append(leagueIDs, int(l.Field(i).Int()))
	}

	for _, leagueId := range leagueIDs {
		if leagueId == service.Config.SofascoreApi.LeaguesID.LaLiga {
			for _, seasonId := range laLigaSeasonIDs {
				t.Logf("Fetching LaLiga SeasonID: %d, LeagueID: %d", seasonId, leagueId)
				if err := service.UpdateLeagueStanding(ctx, seasonId, leagueId); err != nil {
					t.Fatalf("Error updating league standing for LaLiga SeasonID %d: %v", seasonId, err)
				}
			}
		} else if leagueId == service.Config.SofascoreApi.LeaguesID.PremierLeague {
			for _, seasonId := range premierLeagueSeasonIDs {
				t.Logf("Fetching Premier League SeasonID: %d, LeagueID: %d", seasonId, leagueId)
				if err := service.UpdateLeagueStanding(ctx, seasonId, leagueId); err != nil {
					t.Fatalf("Error updating league standing for Premier League SeasonID %d: %v", seasonId, err)
				}
			}
		}
	}

	t.Log("Calling Sofascore API...")
	t.Log("Sofascore API Publish test completed successfully")

}
