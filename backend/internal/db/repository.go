package db

import (
	"time"
)

func SaveUpload(filename, path string) error {
	_, err := DB.Exec(`
		INSERT INTO pdf_uploads (filename, path, upload_time)
		VALUES ($1, $2, $3)
	`, filename, path, time.Now())
	return err
}
