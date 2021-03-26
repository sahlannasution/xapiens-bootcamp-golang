package models

import (
	"time"

	"gorm.io/gorm"
)

// Reviews struct
type Reviews struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Review  string `json:"review"`
	Rate    int    `json:"rate"`
	UsersID uint   `json:"users_id"`
	// Users     []Users        `gorm:"foreignKey:UsersID" json:"users"`
	MoviesID uint   `json:"movies_id"`
	Movies   Movies `gorm:"association_foreignkey:ID" json:"movies"`
	Users    Users  `gorm:"association_foreignkey:ID" json:"users"`
	// Movies    []Movies       `gorm:"foreignKey:MoviesID" json:"movies"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
