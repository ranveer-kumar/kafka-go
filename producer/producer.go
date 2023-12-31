package producer

import (
	"context"
	"kafka-go/internal/config"
	kafka "kafka-go/pkg/kafka"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(cfg config.KafkaConfig) (*Producer, error) {
	client, err := kafka.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	return &Producer{client: client}, nil
}

func (p *Producer) SendMessage(topic, message string) error {
	ctx := context.Background()
	err := p.client.ProduceSync(ctx, &kgo.Record{
		Topic: topic,
		Value: []byte(message),
	}).FirstErr()

	if err != nil {
		log.Printf("Failed to send message to topic %s: %v\n", topic, err)
		return err
	}

	log.Printf("Message sent to topic %s: %s\n", topic, message)
	return nil
}
