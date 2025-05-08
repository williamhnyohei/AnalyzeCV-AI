package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define conexão com o Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "pdf-uploaded",
		GroupID:   "analyze-cv-consumer-group",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer reader.Close()

	fmt.Println("Consumer iniciado e aguardando mensagens...")

	// Loop para consumir mensagens
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			continue
		}
		log.Printf("Mensagem recebida! Key: %s | Value: %s", string(msg.Key), string(msg.Value))
	}
}
