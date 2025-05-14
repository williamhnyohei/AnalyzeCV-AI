package db

import "time"

type PDFUpload struct {
	Filename   string
	Path       string
	UploadTime time.Time
}
