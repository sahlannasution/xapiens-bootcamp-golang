package models

import (
	"time"

	"gorm.io/gorm"
)

// Post struct
type Post struct {
	UserID    uint   `json:"user_id"`
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
