package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"log"
	"storage/config"
	"strconv"
	"strings"
)

type MyKafka struct {
	logger        *zap.SugaredLogger
	settings      config.Settings
	writer        *kafka.Writer
	reader        *kafka.Reader
	detailService DetailService
}

func NewKafka(logger *zap.SugaredLogger, settings config.Settings, detailService DetailService) MyKafka {
	return MyKafka{
		logger:        logger,
		settings:      settings,
		writer:        getKafkaWriter(settings.KafkaUrl, settings.WriteTopic),
		reader:        getKafkaReader(settings.KafkaUrl, settings.ReadTopic, "0"),
		detailService: detailService,
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
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
	err := k.writer.WriteMessages(ctx, msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (k MyKafka) Consume() {
	fmt.Println("start consuming ... !!")

	for {
		msg, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		k.logger.Debug("message at topic:%v partition:%v offset:%v	%s = %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		var str = string(msg.Key)
		var i, _ = strconv.ParseInt(str, 10, 64)
		flag, key, value := k.detailService.Processing(i)
		if flag {
			k.Produce(context.Background(), key, value)
		}
	}
}
