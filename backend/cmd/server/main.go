package main

import (
	"fmt"
	"log"
	"net/http"
	"analyze-cv-ai/internal/handler"
)

func main() {
	http.HandleFunc("/health", handler.HealthCheck)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
