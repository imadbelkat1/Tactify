package config

import (
	"log"
	"path/filepath"
	"runtime"

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
	// Get the directory where this config.go file is located
	_, filename, _, _ := runtime.Caller(0)
	kafkaConfigDir := filepath.Dir(filename)
	kafkaRootDir := filepath.Dir(kafkaConfigDir) // Go up one level to kafka/

	viper.SetConfigFile(filepath.Join(kafkaRootDir, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Kafka: Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return &config
}
