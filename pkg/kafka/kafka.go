package kafka

import (
	"context"

	"github.com/AyanChoudhary/transponster/pkg/utils"
)

// Kafka type to hold reader and writer instances
type Kafka struct {
	writers []*Writer
	reader  *Reader
}

// GenerateFromConfig generates readers and writes from a given config file
func GenerateFromConfig(ctx context.Context, configPath string) *Kafka {
	config := utils.GenerateConfigReader("./config.json")
	writerTopic := config.String("writer.topic")
	writerBrokers := config.Strings("writer.brokers")

	writers := []*Writer{GenerateWriter(ctx, writerBrokers, writerTopic)}

	readerTopic := config.String("reader.topic")
	readerGroupID := config.String("reader.groupID")
	readerBrokers := config.Strings("reader.Brokers")
	reader := GenerateReader(ctx, readerBrokers, readerTopic, readerGroupID)

	return &Kafka{
		writers,
		reader,
	}
}
