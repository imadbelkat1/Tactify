package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	//kafka
	KafkaTopics []string `mapstructure:"KAFKA_TOPICS"`

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
	viper.SetConfigFile(".env")
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
