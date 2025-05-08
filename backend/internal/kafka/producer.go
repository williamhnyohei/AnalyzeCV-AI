package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaProducer() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "pdf-uploaded",
		Balancer: &kafka.LeastBytes{},
	})
	fmt.Println("Kafka Producer iniciado com sucesso")
}

func SendMessage(filename string) error {
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("pdf-%d", time.Now().UnixNano())),
		Value: []byte(filename),
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Erro ao enviar mensagem para o Kafka: %v", err)
		return err
	}

	log.Printf("Mensagem enviada para o Kafka: %s", filename)
	return nil
}

func CloseProducer() {
	if writer != nil {
		writer.Close()
	}
}
