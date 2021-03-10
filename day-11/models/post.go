package models

import (
	"time"

	"gorm.io/gorm"
)

// Post struct
type Post struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
