package db

import "time"

type PDF struct {
	ID        int
	Filename  string
	Status    string
	CreatedAt time.Time
}
