package main

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/imadbelkat1/fpl-service/config"
	fplApi "github.com/imadbelkat1/fpl-service/internal/api"

	"github.com/imadbelkat1/fpl-service/internal/models"
	teamService "github.com/imadbelkat1/fpl-service/internal/services"

	kafka "github.com/imadbelkat1/kafka"
	kafkaConfig "github.com/imadbelkat1/kafka/config"
)

func TestTeamApiService_RealAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping real API test")
	}

	// Setup service with real client
	service := &teamService.TeamApiService{
		Client: fplApi.NewFplApiClient(),
	}

	// Start consumer to verify messages
	go consumeTeamMessages(t)

	// Wait for consumer to start
	time.Sleep(200 * time.Millisecond)

	// Test with real API
	log.Println("Calling real FPL API...")
	err := service.UpdateTeams()
	if err != nil {
		t.Fatalf("UpdateTeams with real API failed: %v", err)
	}

	// Wait for messages to be processed
	time.Sleep(3 * time.Second)

	t.Log("Real API test completed successfully")
}

func consumeTeamMessages(t *testing.T) {
	cfg := config.LoadConfig()
	kafkaCfg := kafkaConfig.LoadConfig()

	consumer := kafka.NewConsumer(
		kafkaCfg,
		cfg.FplTeamsTopic,
	)
	defer consumer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messages, errors := consumer.Subscribe(ctx)

	messageCount := 0
	for {
		select {
		case msg := <-messages:
			messageCount++

			t.Logf("Received message %d: key=%s, value=%s", messageCount, string(msg.Key), string(msg.Value))

			var team models.Team
			if err := json.Unmarshal(msg.Value, &team); err != nil {
				t.Errorf("Failed to unmarshal team: %v", err)
				continue
			}

			t.Logf("Received Team %d: %s (%s)",
				team.ID, team.Name, team.ShortName)

			// FPL has 20 teams
			if messageCount >= 20 {
				t.Logf("Successfully received all %d teams from real API", messageCount)
				return
			}

		case err := <-errors:
			if err != nil {
				t.Logf("Consumer error: %v", err)
			}

		case <-ctx.Done():
			t.Logf("Test finished with %d teams received", messageCount)
			return
		}
	}
}
