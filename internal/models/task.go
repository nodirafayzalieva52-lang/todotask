package models

import(
	"time"
) 

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"user_id"`
	Title       string    `gorm:"size:100" json:"title"`
	Description string    `gorm:"size:100" json:"description"`
	Status      string    `gorm:"size:100" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
