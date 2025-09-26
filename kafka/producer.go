package kafka

import (
	"context"
	"time"

	"github.com/imadbelkat1/kafka/config"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	kafkaCfg := config.LoadConfig()

	return &Producer{
		writer: &kafka.Writer{
			Addr:            kafka.TCP(kafkaCfg.KafkaBroker),
			RequiredAcks:    kafka.RequiredAcks(parseAcks(kafkaCfg.KafkaAcks)),
			Async:           false,
			Compression:     parseCompression(kafkaCfg.KafkaCompressionType),
			BatchSize:       kafkaCfg.KafkaBatchSize,
			BatchTimeout:    time.Duration(kafkaCfg.KafkaLingerMs) * time.Millisecond,
			WriteBackoffMin: time.Duration(kafkaCfg.KafkaRetryBackoffMs) * time.Millisecond,
			WriteTimeout:    time.Duration(kafkaCfg.KafkaDeliveryTimeoutMs) * time.Millisecond,
		},
	}
}

func (p *Producer) Publish(ctx context.Context, topic string, key, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

// Helper functions to add
func parseAcks(acks string) int {
	switch acks {
	case "all", "-1":
		return -1
	case "1":
		return 1
	case "0":
		return 0
	default:
		return -1 // default to "all"
	}
}

func parseCompression(compression string) kafka.Compression {
	switch compression {
	case "snappy":
		return kafka.Snappy
	case "gzip":
		return kafka.Gzip
	case "lz4":
		return kafka.Lz4
	case "zstd":
		return kafka.Zstd
	default:
		return kafka.Snappy // default
	}
}
