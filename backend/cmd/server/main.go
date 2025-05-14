package main

import (
	"analyze-cv-ai/internal/db"
	"analyze-cv-ai/internal/handler"
	"analyze-cv-ai/internal/kafka"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.Init()
	kafka.InitKafkaProducer()
	defer kafka.CloseProducer()

	http.HandleFunc("/health", handler.HealthCheck)
	http.HandleFunc("/upload", handler.UploadPDF)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
