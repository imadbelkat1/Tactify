package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	teamOverallStatsService "github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestTeamOverallStats(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	service := &teamOverallStatsService.TeamOverallStatsService{
		Config:   config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
		Standing: &teamOverallStatsService.LeagueStandingService{
			Config:   config.LoadConfig(),
			Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
			Producer: kafka.NewProducer(),
		},
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
				if err := service.UpdateAllTeamsOverallStats(ctx, leagueId, seasonId); err != nil {
					t.Fatalf("Error updating teams overall stats for La Liga SeasonID %d: %v", seasonId, err)
				}
			}
		} else if leagueId == service.Config.SofascoreApi.LeaguesID.PremierLeague {
			for _, seasonId := range premierLeagueSeasonIDs {
				if err := service.UpdateAllTeamsOverallStats(ctx, leagueId, seasonId); err != nil {
					t.Fatalf("Error updating teams overall stats for Premier League SeasonID %d: %v", seasonId, err)
				}
			}
		}
	}

	t.Log("Calling Sofascore API...")
	t.Log("Sofascore API Publish test completed successfully")

}
