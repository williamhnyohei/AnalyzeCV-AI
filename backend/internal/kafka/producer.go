package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	WriteMessages(ctx context.Context, msgs ...kafkago.Message) error
	Close() error
}

var writer KafkaProducer

func InitKafkaProducer() {
	writer = kafkago.NewWriter(kafkago.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "pdf-uploaded",
		Balancer: &kafkago.LeastBytes{},
	})
	fmt.Println("Kafka Producer iniciado com sucesso")
}

func SetWriter(w KafkaProducer) {
	writer = w
}

func SendMessage(filename string) error {
	if writer == nil {
		return fmt.Errorf("kafka writer não inicializado")
	}

	msg := kafkago.Message{
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
