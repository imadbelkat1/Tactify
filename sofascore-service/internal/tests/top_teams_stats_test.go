package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
	topTeamsStatservice "github.com/imadbelkat1/sofascore-service/internal/services"
)

func TestTopTeamsStatsService(t *testing.T) {
	ctx := context.Background()
	if testing.Short() {
		t.Skip("Skipping API test")
	}

	service := &topTeamsStatservice.TopTeamsStatsService{
		Config:   *config.LoadConfig(),
		Client:   sofascore_api.NewSofascoreApiClient(config.LoadConfig()),
		Producer: kafka.NewProducer(),
	}

	log.Println("Calling FPL API...")

	seasonId := service.Config.SofascoreApi.SeasonsIDs.PremierLeague2526
	leagueId := service.Config.SofascoreApi.LeaguesIDs.PremierLeague
	log.Println(seasonId)
	log.Println(leagueId)

	start := time.Now()
	err := service.UpdateLeagueTopTeamsStats(ctx, seasonId, leagueId)
	if err != nil {
		t.Fatalf("GetTopTeamsStats failed: %v", err)
	}
	elapsed := time.Since(start)

	log.Printf("Publishing completed in: %v", elapsed)
	t.Log("Sofascore API test completed successfully")
}
