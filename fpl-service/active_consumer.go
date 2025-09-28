package main

import (
	"context"
	"fmt"

	"github.com/imadbelkat1/fpl-service/config"
	"github.com/imadbelkat1/kafka"
	kafkaConfig "github.com/imadbelkat1/kafka/config"
)

func main() {
	cfg := config.LoadConfig()
	kafkaCfg := kafkaConfig.LoadConfig()

	// Get all available topics
	topics := []string{
		cfg.TopicsName.FplTeams,
		cfg.TopicsName.FplFixtures,
		cfg.TopicsName.FplPlayerMatchStats,
		cfg.TopicsName.FplLiveEvent,
		// Add other topics here as needed
	}

	for _, topic := range topics {
		go func(topicName string) {
			consumer := kafka.NewConsumer(
				kafkaCfg,
				topicName,
			)
			defer func(consumer *kafka.Consumer) {
				err := consumer.Close()
				if err != nil {
					fmt.Printf("Error closing consumer for topic %s: %v\n", topicName, err)
				}
			}(consumer)

			ctx := context.Background()
			messages, errors := consumer.Subscribe(ctx)

			fmt.Printf("Starting to listen on topic: %s\n", topicName)

			for {
				select {
				case msg := <-messages:
					fmt.Printf("Topic [%s] - Received message: key=%s, value=%s\n",
						topicName, string(msg.Key), string(msg.Value))

				case err := <-errors:
					if err != nil {
						fmt.Printf("Topic [%s] - Consumer error: %v\n", topicName, err)
					}

				case <-ctx.Done():
					fmt.Printf("Topic [%s] - Consumer stopped\n", topicName)
					return
				}
			}
		}(topic) // Adjust index as needed
	}

	// Keep program running to listen for messages
	select {}
}
