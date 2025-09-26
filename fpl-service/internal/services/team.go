package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imadbelkat1/fpl-service/config"
	"github.com/imadbelkat1/fpl-service/internal/api"
	"github.com/imadbelkat1/fpl-service/internal/models"

	teamProducer "github.com/imadbelkat1/kafka"
)

type TeamApiService struct {
	Client *fpl_api.FplApiClient
}

func (s *TeamApiService) UpdateTeams() error {
	var bootstrap models.BootstrapResponse
	producer := teamProducer.NewProducer()
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
		return err
	}

	for _, t := range bootstrap.Teams {
		teamJSON, err := json.Marshal(t)
		if err != nil {
			return fmt.Errorf("marshaling team: %w", err)
		}
		err = producer.Publish(ctx, cfg.FplTeamsTopic, []byte(fmt.Sprintf("%d", t.ID)), teamJSON)
		if err != nil {
			return fmt.Errorf("publishing team to Kafka error: %w", err)
		}
	}

	return nil
}

func (s *TeamApiService) GetTeam(id int) (*models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, bootstrap); err != nil {
		return nil, err
	}

	for _, t := range bootstrap.Teams {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("team with id %d not found", id)
}

func (s *TeamApiService) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	var bootstrap models.BootstrapResponse

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
		return nil, err
	}

	return bootstrap.Teams, nil
}

func (s *TeamApiService) GetTeamsByStrength(minStrength int) ([]models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
		return nil, err
	}

	var filteredTeams []models.Team
	for _, t := range bootstrap.Teams {
		if t.Strength >= minStrength {
			filteredTeams = append(filteredTeams, t)
		}
	}

	return filteredTeams, nil
}
