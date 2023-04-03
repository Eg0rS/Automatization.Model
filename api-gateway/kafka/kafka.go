package kafka

import (
	"api-gateway/config"
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"log"
)

type MyKafka struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	kafka    *kafka.Writer
}

func NewKafka(logger *zap.SugaredLogger, settings config.Settings) MyKafka {
	return MyKafka{
		logger:   logger,
		settings: settings,
		kafka:    getKafkaWriter(settings.KafkaUrl, settings.WriteTopic),
	}
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (k MyKafka) Produce(ctx context.Context, key []byte, value []byte) {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	err := k.kafka.WriteMessages(ctx, msg)

	if err != nil {
		log.Fatalln(err)
	}
}
