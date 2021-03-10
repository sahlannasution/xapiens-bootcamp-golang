package models

import (
	"time"

	"gorm.io/gorm"
)

// Todos struct
type Todos struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
