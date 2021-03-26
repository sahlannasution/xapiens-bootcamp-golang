package models

import (
	"time"

	"gorm.io/gorm"
)

// MovieGenres struct
type MovieGenres struct {
	ID       uint `gorm:"primarykey" json:"id"`
	MoviesID uint `json:"movies_id"`
	GenresID uint `json:"genres_id"`
	// Genres    []Genres       `gorm:"foreignKey:ID" json:"genres"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
