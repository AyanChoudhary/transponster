package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// Reader is wrapper around kafka.Reader
type Reader struct {
	reader *kafka.Reader
}

// GenerateReader creates a reader kafka instance to respond to incoming requests
func GenerateReader(ctx context.Context, brokers []string, topic string, groupID string) *Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	return &Reader{reader}
}
