package storage

import (
	"io"
	"os"
	"path/filepath"
)

const UploadDir = "/tmp/uploads"

func SaveFile(file io.Reader, filename string) (string, error) {
	// Garante que o diretório existe
	err := os.MkdirAll(UploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Cria o arquivo no sistema
	filePath := filepath.Join(UploadDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copia o conteúdo para o novo arquivo
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
