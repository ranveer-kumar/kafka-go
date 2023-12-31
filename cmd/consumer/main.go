package main

import (
	"kafka-go/consumer"
	"kafka-go/internal/config"

	"log"
)

func main() {

	kafkaConfig := config.LoadKafkaConfig()

	// Create a Kafka consumer with 4 worker Goroutines for message processing
	consumer, err := consumer.NewConsumer(kafkaConfig, 4)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Start consuming messages
	consumer.StartConsuming()

}
