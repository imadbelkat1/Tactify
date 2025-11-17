package services

import (
	"context"

	"github.com/imadbelkat1/kafka"
	"github.com/imadbelkat1/sofascore-service/config"
	sofascore_api "github.com/imadbelkat1/sofascore-service/internal/api"
)

type PlayersService struct {
	Config   config.SofascoreConfig
	Client   *sofascore_api.SofascoreApiClient
	Producer *kafka.Producer
}

// get unique tour id -> fetch its standing -> get team ids from it -> fetch players

func (p *PlayersService) GetPlayerInfo(ctx context.Context, playerId int) error {
	return nil
}
