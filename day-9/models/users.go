package models

import "gorm.io/gorm"

// Users struct
type Users struct {
	ID       uint   `gorm:"primarykey, autoIncrement" json:"id"`
	Email    string `gorm:"primarykey" json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	gorm.Model
}

// GetUsers struct
type GetUsers struct {
	ID       string `gorm:"primarykey" json:"id"`
	Email    string `gorm:"primarykey" json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	gorm.Model
}
