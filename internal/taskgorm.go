package models

import (
	"time"
)

type Task struct {
	ID uint
	Title string
	Description string
	Status string
	CreatedAt time.Time
}

