package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imadbelkat1/fpl-service/internal/api"
	"github.com/imadbelkat1/fpl-service/internal/models"
)

type TeamApiService struct {
	Client *fpl_api.FplApiClient
}

func (s *TeamApiService) getBootstrapData(ctx context.Context) (*models.BootstrapResponse, error) {
	var bootstrap models.BootstrapResponse
	if err := s.Client.GetAndUnmarshal(ctx, bootstrapEndpoint, &bootstrap); err != nil {
		return nil, err
	}
	return &bootstrap, nil
}

func (s *TeamApiService) UpdateTeams() error {
	ctx := context.Background()
	bootstrap, err := s.getBootstrapData(ctx)

	if err = s.Client.GetAndUnmarshal(ctx, bootstrapEndpoint, &bootstrap); err != nil {
		return err
	}

	for _, t := range bootstrap.Teams {
		teamJSON, err := json.Marshal(t)
		if err != nil {
			return fmt.Errorf("marshaling team: %w", err)
		}
		err = Publish(ctx, teamProducer, teamsTopic, []byte(fmt.Sprintf("%d", t.ID)), teamJSON)
		if err != nil {
			return fmt.Errorf("publishing team to Kafka error: %w", err)
		}
	}

	return nil
}

func (s *TeamApiService) GetTeam(id int) (*models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	if err := s.Client.GetAndUnmarshal(ctx, bootstrapEndpoint, bootstrap); err != nil {
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

	if err := s.Client.GetAndUnmarshal(ctx, bootstrapEndpoint, &bootstrap); err != nil {
		return nil, err
	}

	return bootstrap.Teams, nil
}

func (s *TeamApiService) GetTeamsByStrength(minStrength int) ([]models.Team, error) {
	var bootstrap models.BootstrapResponse
	ctx := context.Background()

	if err := s.Client.GetAndUnmarshal(ctx, bootstrapEndpoint, &bootstrap); err != nil {
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
