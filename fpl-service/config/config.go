package config

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	FplApi           FplApi
	TopicsName       TopicsName
	Topics           TopicsRetention
	ConsumersGroupID ConsumersGroupID
}

type FplApi struct {
	//fpl api
	BaseUrl               string `mapstructure:"FPLAPI_BASE_URL"`
	Bootstrap             string `mapstructure:"FPLAPI_BOOTSTRAP"`
	Fixtures              string `mapstructure:"FPLAPI_FIXTURES"`
	PlayerSummary         string `mapstructure:"FPLAPI_PLAYER_SUMMARY"`
	Entry                 string `mapstructure:"FPLAPI_ENTRY"`
	EntryEvent            string `mapstructure:"FPLAPI_ENTRY_EVENT"`
	EntryHistory          string `mapstructure:"FPLAPI_ENTRY_HISTORY"`
	EntryTransfers        string `mapstructure:"FPLAPI_ENTRY_TRANSFERS"`
	EntryPicks            string `mapstructure:"FPLAPI_ENTRY_PICKS"`
	LiveEvent             string `mapstructure:"FPLAPI_LIVE_EVENT"`
	LeagueClassicStanding string `mapstructure:"FPLAPI_LEAGUE_CLASSIC_STANDING"`
	LeagueH2hStanding     string `mapstructure:"FPLAPI_LEAGUE_H2H_STANDING"`
}
type TopicsName struct {
	// FPL Core Data Topics
	FplPlayers               string `mapstructure:"TOPICSNAME_FPL_PLAYERS"`
	FplTeams                 string `mapstructure:"TOPICSNAME_FPL_TEAMS"`
	FplFixtures              string `mapstructure:"TOPICSNAME_FPL_FIXTURES"`
	FplFixtureDetails        string `mapstructure:"TOPICSNAME_FPL_FIXTURE_DETAILS"`
	FplPlayerMatchStats      string `mapstructure:"TOPICSNAME_FPL_PLAYER_MATCH_STATS"`
	FplLiveEvent             string `mapstructure:"TOPICSNAME_FPL_LIVE_EVENT"`
	FplEntry                 string `mapstructure:"TOPICSNAME_FPL_ENTRY"`
	FplEntryEvent            string `mapstructure:"TOPICSNAME_FPL_ENTRY_EVENT"`
	FplEntryHistory          string `mapstructure:"TOPICSNAME_FPL_ENTRY_HISTORY"`
	FplEntryTransfers        string `mapstructure:"TOPICSNAME_FPL_ENTRY_TRANSFERS"`
	FplEntryPicks            string `mapstructure:"TOPICSNAME_FPL_ENTRY_PICKS"`
	FplLeagueClassicStanding string `mapstructure:"TOPICSNAME_FPL_LEAGUE_CLASSIC_STANDING"`
	FplLeagueH2hStanding     string `mapstructure:"TOPICSNAME_FPL_LEAGUE_H2H_STANDING"`
}
type TopicsRetention struct {
	FplPlayers               string `mapstructure:"TOPICSRETENTION_FPL_PLAYERS"`
	FplTeams                 string `mapstructure:"TOPICSRETENTION_FPL_TEAMS"`
	FplFixtures              string `mapstructure:"TOPICSRETENTION_FPL_FIXTURES"`
	FplPlayerMatchStats      string `mapstructure:"TOPICSRETENTION_FPL_PLAYER_MATCH_STATS"`
	FplEntry                 string `mapstructure:"TOPICSRETENTION_FPL_ENTRY"`
	FplEntryEvent            string `mapstructure:"TOPICSRETENTION_FPL_ENTRY_EVENT"`
	FplEntryHistory          string `mapstructure:"TOPICSRETENTION_FPL_ENTRY_HISTORY"`
	FplEntryTransfers        string `mapstructure:"TOPICSRETENTION_FPL_ENTRY_TRANSFERS"`
	FplEntryPicks            string `mapstructure:"TOPICSRETENTION_FPL_ENTRY_PICKS"`
	FplLeagueClassicStanding string `mapstructure:"TOPICSRETENTION_FPL_LEAGUE_CLASSIC_STANDING"`
	FplLeagueH2hStanding     string `mapstructure:"TOPICSRETENTION_FPL_LEAGUE_H2H_STANDING"`
}
type ConsumersGroupID struct {
	// Consumer Group IDs
	Teams                  string `mapstructure:"CONSUMERSGROUPID_KAFKA_TEAMS"`
	Fixtures               string `mapstructure:"CONSUMERSGROUPID_KAFKA_FIXTURES"`
	Players                string `mapstructure:"CONSUMERSGROUPID_KAFKA_PLAYERS"`
	PlayersStats           string `mapstructure:"CONSUMERSGROUPID_KAFKA_PLAYERS_STATS"`
	Live                   string `mapstructure:"CONSUMERSGROUPID_KAFKA_LIVE_EVENT"`
	Entries                string `mapstructure:"CONSUMERSGROUPID_KAFKA_ENTRY"`
	EntriesEvent           string `mapstructure:"CONSUMERSGROUPID_KAFKA_ENTRY_EVENT"`
	EntriesHistory         string `mapstructure:"CONSUMERSGROUPID_KAFKA_ENTRY_HISTORY"`
	EntriesTransfers       string `mapstructure:"CONSUMERSGROUPID_KAFKA_ENTRY_TRANSFERS"`
	EntriesPicks           string `mapstructure:"CONSUMERSGROUPID_KAFKA_ENTRY_PICKS"`
	LeaguesClassicStanding string `mapstructure:"CONSUMERSGROUPID_KAFKA_LEAGUES_CLASSIC_STANDING"`
	LeaguesH2hStanding     string `mapstructure:"CONSUMERSGROUPID_KAFKA_LEAGUES_H2H_STANDING"`
	Test                   string `mapstructure:"CONSUMERSGROUPID_KAFKA_TEST"`
}

func LoadConfig() *Config {
	_, filename, _, _ := runtime.Caller(0)
	ConfigDir := filepath.Dir(filename)
	RootDir := filepath.Dir(ConfigDir)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(RootDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fpl-service: Error reading config file, %s", err)
	}

	config := &Config{}
	mapViperToStruct(config)
	return config
}

func mapViperToStruct(v interface{}) {
	val := reflect.ValueOf(v).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Struct && field.CanSet() {
			mapStructFields(field)
		}
	}
}

func mapStructFields(structField reflect.Value) {
	for i := 0; i < structField.NumField(); i++ {
		field := structField.Field(i)
		if tag := structField.Type().Field(i).Tag.Get("mapstructure"); tag != "" && field.CanSet() {
			field.SetString(viper.GetString(tag))
		}
	}
}
