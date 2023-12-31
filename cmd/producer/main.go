package main

import (
	"fmt"
	"kafka-go/internal/config"
	"kafka-go/producer"
	"log"
	"net/http"
)

func main() {
	kafkaCfg := config.LoadKafkaConfig()

	prod, err := producer.NewProducer(kafkaCfg)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	http.HandleFunc("/send", producer.SendMessageHandler(prod))
	port := 8080 
	log.Printf("Server listening on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
