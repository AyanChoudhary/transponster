package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

//Writer is a wrapper around kafka.Writer
type Writer struct {
	writer *kafka.Writer
}

// GenerateWriter creates a writer kafka instance to make outgoing responses
func GenerateWriter(ctx context.Context, brokers []string, topic string) *Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return &Writer{writer}
}
