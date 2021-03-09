package models

import (
	"time"

	"gorm.io/gorm"
)

// Todos struct
type Todos struct {
	UserID    uint   `json:"user_id"`
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
