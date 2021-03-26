package models

import (
	"time"

	"gorm.io/gorm"
)

// Users struct
type Users struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	// Reviews   []Reviews      `gorm:"foreignKey:UsersID" json:"reviews"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
