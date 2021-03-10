package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment struct
type Comment struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	PostID    uint   `json:"postId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
