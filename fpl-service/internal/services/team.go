package services

import (
	"context"
	"fmt"
	"tactify/fpl-service/config"
	"tactify/fpl-service/internal/api"
	"tactify/fpl-service/internal/models"

	producer "tactify/kafka/internal/producer"
)

type TeamApiService struct {
	client *fpl_api.FplApiClient
}

func (s *TeamApiService) UpdateTeams() error {
	var bootstrap models.BootstrapResponse
	producer = producer.NewProducer()
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
		return err
	}

	for _, t := range bootstrap.Teams {
	}

	return nil
}

func (s *TeamApiService) GetTeam(id int) (*models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.client.GetAndUnmarshal(ctx, endpoint, bootstrap); err != nil {
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

	if err := s.client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
		return nil, err
	}

	return bootstrap.Teams, nil
}

func (s *TeamApiService) GetTeamsByStrength(minStrength int) ([]models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	cfg := config.LoadConfig()
	endpoint := cfg.FplApiBootstrap

	if err := s.client.GetAndUnmarshal(ctx, endpoint, &bootstrap); err != nil {
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
