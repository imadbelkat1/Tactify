package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// FPL Core Data Topics
	FplPlayersTopic  string `mapstructure:"FPL_PLAYERS_TOPIC"`
	FplTeamsTopic    string `mapstructure:"FPL_TEAMS_TOPIC"`
	FplFixturesTopic string `mapstructure:"FPL_FIXTURES_TOPIC"`

	// FPL Entry Topics
	FplEntryTopic          string `mapstructure:"FPL_ENTRY_TOPIC"`
	FplEntryEventTopic     string `mapstructure:"FPL_ENTRY_EVENT_TOPIC"`
	FplEntryHistoryTopic   string `mapstructure:"FPL_ENTRY_HISTORY_TOPIC"`
	FplEntryTransfersTopic string `mapstructure:"FPL_ENTRY_TRANSFERS_TOPIC"`
	FplEntryPicksTopic     string `mapstructure:"FPL_ENTRY_PICKS_TOPIC"`

	// FPL League Topics
	FplLeagueClassicStandingTopic string `mapstructure:"FPL_LEAGUE_CLASSIC_STANDING_TOPIC"`
	FplLeagueH2hStandingTopic     string `mapstructure:"FPL_LEAGUE_H2H_STANDING_TOPIC"`

	//fpl api
	FplApiBaseUrl               string `mapstructure:"FPL_API_BASE_URL"`
	FplApiBootstrap             string `mapstructure:"FPL_API_BOOTSTRAP"`
	FplApiFixtures              string `mapstructure:"FPL_API_FIXTURES"`
	FplApiPlayerSummary         string `mapstructure:"FPL_API_PLAYER_SUMMARY"`
	FplApiEntry                 string `mapstructure:"FPL_API_ENTRY"`
	FplApiEntryEvent            string `mapstructure:"FPL_API_ENTRY_EVENT"`
	FplApiEntryHistory          string `mapstructure:"FPL_API_ENTRY_HISTORY"`
	FplApiEntryTransfers        string `mapstructure:"FPL_API_ENTRY_TRANSFERS"`
	FplApiEntryPicks            string `mapstructure:"FPL_API_ENTRY_PICKS"`
	FplApiLiveEvent             string `mapstructure:"FPL_API_LIVE_EVENT"`
	FplApiLeagueClassicStanding string `mapstructure:"FPL_API_LEAGUE_CLASSIC_STANDING"`
	FplApiLeagueH2hStanding     string `mapstructure:"FPL_API_LEAGUE_H2H_STANDING"`
}

func LoadConfig() *Config {
	viper.SetConfigFile("../.env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return &config
}
