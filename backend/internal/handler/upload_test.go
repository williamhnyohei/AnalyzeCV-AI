package handler

import (
	"analyze-cv-ai/internal/kafka"
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	segmentioKafka "github.com/segmentio/kafka-go"
)

type mockKafkaWriter struct{}

func (m *mockKafkaWriter) WriteMessages(ctx context.Context, msgs ...segmentioKafka.Message) error {
	return nil
}

func (m *mockKafkaWriter) Close() error {
	return nil
}

func TestUploadPDF(t *testing.T) {
	kafka.SetWriter(&mockKafkaWriter{})

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("conteudo falso de PDF"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	UploadPDF(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status OK, mas recebeu %v", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Arquivo salvo com sucesso") {
		t.Errorf("esperado resposta de sucesso, mas recebeu: %v", rr.Body.String())
	}
}
