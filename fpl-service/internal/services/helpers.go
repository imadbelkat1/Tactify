package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imadbelkat1/fpl-service/config"
	Producer "github.com/imadbelkat1/kafka"
)

type ProcessedModel struct {
	ID   int
	Data []byte
}

var cfg = config.LoadConfig()

var bootstrapEndpoint = cfg.FplApi.Bootstrap                         // /bootstrap-static/
var fixturesEndpoint = cfg.FplApi.Fixtures                           // /fixtures/
var playerSummaryEndpoint = cfg.FplApi.PlayerSummary                 // /element-summary/%d/
var entryEndpoint = cfg.FplApi.Entry                                 // /entry/%d/
var entryEventEndpoint = cfg.FplApi.EntryEvent                       // /entry/%d/event/%d/
var entryHistoryEndpoint = cfg.FplApi.EntryHistory                   // /entry/%d/history/
var entryTransfersEndpoint = cfg.FplApi.EntryTransfers               // /entry/%d/transfers/
var entryPicksEndpoint = cfg.FplApi.EntryPicks                       // /entry/%d/event/%d/picks/
var leagueClassicStandingEndpoint = cfg.FplApi.LeagueClassicStanding // /leagues-classic/%d/standings/
var leagueH2hStandingEndpoint = cfg.FplApi.LeagueH2hStanding         // /leagues-h2h/%d/standings/
var liveEventEndpoint = cfg.FplApi.LiveEvent                         // /event/%d/live/

var fixturesTopic = cfg.TopicsName.FplFixtures
var fixtureDetailsTopic = cfg.TopicsName.FplFixtureDetails
var teamsTopic = cfg.TopicsName.FplTeams
var playersTopic = cfg.TopicsName.FplPlayers
var playerMatchStatsTopic = cfg.TopicsName.FplPlayerMatchStats
var entryTopic = cfg.TopicsName.FplEntry
var entryEventTopic = cfg.TopicsName.FplEntryEvent
var entryHistoryTopic = cfg.TopicsName.FplEntryHistory
var entryTransfersTopic = cfg.TopicsName.FplEntryTransfers
var entryPicksTopic = cfg.TopicsName.FplEntryPicks
var leagueClassicTopic = cfg.TopicsName.FplLeagueClassicStanding
var leagueH2hTopic = cfg.TopicsName.FplLeagueH2hStanding
var liveEventTopic = cfg.TopicsName.FplLiveEvent

var teamProducer = Producer.NewProducer()
var fixtureProducer = Producer.NewProducer()
var statsProducer = Producer.NewProducer()
var playerProducer = Producer.NewProducer()
var entryProducer = Producer.NewProducer()
var entryEventProducer = Producer.NewProducer()
var entryHistoryProducer = Producer.NewProducer()
var entryTransfersProducer = Producer.NewProducer()
var entryPicksProducer = Producer.NewProducer()
var leagueClassicProducer = Producer.NewProducer()
var leagueH2hProducer = Producer.NewProducer()
var liveEventProducer = Producer.NewProducer()

func deleteKey(T any, key []string) (map[string]interface{}, error) {
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

	for _, k := range key {
		delete(data, k)
	}

	return data, nil
}

func processDelete(model interface{}, toBeDeleted []string) ([]byte, error) {
	newElement, err := deleteKey(model, toBeDeleted)
	if err != nil {
		fmt.Errorf("failed to delete keys from newElement: %v", err)
		return nil, err
	}

	elementJSON, err := json.Marshal(newElement)
	if err != nil {
		fmt.Errorf("failed to marshal elementJSON: %v", err)
		return nil, err
	}

	return elementJSON, nil
}

func Publish(ctx context.Context, producer *Producer.Producer, topic string, key []byte, value []byte) error {
	err := producer.Publish(ctx, topic, key, value)
	if err != nil {
		return fmt.Errorf("publishing to Kafka error: %w", err)
	}
	return nil
}
