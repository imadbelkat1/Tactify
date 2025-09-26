package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	KafkaBroker            string `mapstructure:"KAFKA_BROKER"`
	KafkaAcks              string `mapstructure:"KAFKA_ACKS"`
	KafkaRetries           int    `mapstructure:"KAFKA_RETRIES"`
	KafkaRetryBackoffMs    int    `mapstructure:"KAFKA_RETRY_BACKOFF_MS"`
	KafkaDeliveryTimeoutMs int    `mapstructure:"KAFKA_DELIVERY_TIMEOUT_MS"`
	KafkaBatchSize         int    `mapstructure:"KAFKA_BATCH_SIZE"`
	KafkaLingerMs          int    `mapstructure:"KAFKA_LINGER_MS"`
	KafkaCompressionType   string `mapstructure:"KAFKA_COMPRESSION_TYPE"`
	KafkaBufferMemory      int    `mapstructure:"KAFKA_BUFFER_MEMORY"`
	KafkaPartitions        int    `mapstructure:"KAFKA_PARTITIONS"`
	KafkaReplication       int    `mapstructure:"KAFKA_REPLICATION"`
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
