package main

import (
	"analyze-cv-ai/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler.HealthCheck)
	http.HandleFunc("/upload", handler.UploadPDF)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
