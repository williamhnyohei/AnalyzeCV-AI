package handler

import (
	"analyze-cv-ai/internal/storage"
	"fmt"
	"net/http"
	"path/filepath"
)

func UploadPDF(w http.ResponseWriter, r *http.Request) {
	// Limita tamanho do upload para 10 MB
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao ler arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Verifica se o arquivo é PDF
	if filepath.Ext(handler.Filename) != ".pdf" {
		http.Error(w, "Arquivo inválido. Apenas PDFs são aceitos.", http.StatusBadRequest)
		return
	}

	// Salvar arquivo usando storage
	filePath, err := storage.SaveFile(file, handler.Filename)
	if err != nil {
		http.Error(w, "Erro ao salvar arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Arquivo salvo com sucesso: %s\n", filePath)
}
