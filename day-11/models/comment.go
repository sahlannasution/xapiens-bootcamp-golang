package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment struct
type Comment struct {
	PostID    uint   `json:"post_id"`
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
