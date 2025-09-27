package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imadbelkat1/fpl-service/config"
	fpl_api "github.com/imadbelkat1/fpl-service/internal/api"
	"github.com/imadbelkat1/fpl-service/internal/models"

	Producer "github.com/imadbelkat1/kafka"
)

type FixturesApiService struct {
	Client *fpl_api.FplApiClient
}

func (s *FixturesApiService) UpdateFixtures() error {
	var fixtures models.Fixtures
	fixtureProducer := Producer.NewProducer()
	statsProducer := Producer.NewProducer()

	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiFixtures

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, &fixtures); err != nil {
		return err
	}

	for _, f := range fixtures {
		// Separate stats from fixture before marshaling
		fixtureStatsJSON, err := json.Marshal(f.Stats)

		newFixture, err := deleteKey(f, "stats")
		if err != nil {
			return fmt.Errorf("failed to delete stats key from fixture with ID: %d: %v", f.ID, err)
		}

		// Marshal the modified fixture without stats
		fixtureJSON, err := json.Marshal(newFixture)

		err = fixtureProducer.Publish(ctx, cfg.FplFixturesTopic, []byte(fmt.Sprintf("%d", f.ID)), fixtureJSON)
		if err != nil {
			return fmt.Errorf("failed to publish fixture with ID: %d to Kafka: %v", f.ID, err)
		}

		err = statsProducer.Publish(ctx, cfg.FplPlayerMatchStatsTopic, []byte(fmt.Sprintf("%d", f.ID)), fixtureStatsJSON)
		if err != nil {
			return fmt.Errorf("failed to publish fixture stats for fixture with ID: %d to Kafka: %v", f.ID, err)
		}
	}
	return nil
}

func deleteKey(T any, key string) (map[string]interface{}, error) {
	Bytes, err := json.Marshal(T)
	if err != nil {
		fmt.Println("Error marshaling struct:", err)
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(Bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshaling to map:", err)
		return nil, err
	}

	delete(data, key)

	return data, nil
}
