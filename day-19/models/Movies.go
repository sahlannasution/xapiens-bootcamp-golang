package models

import (
	"time"

	"gorm.io/gorm"
)

// Movies struct
type Movies struct {
	ID          uint          `gorm:"primarykey" json:"id"`
	Title       string        `json:"title"`
	Year        int           `json:"year"`
	Ratings     int           `json:"ratings"`
	MovieGenres []MovieGenres `gorm:"foreignKey:MoviesID" json:"movie_genre"`
	Genres      []Genres      `gorm:"many2many:movie_genres" json:"genres"`
	// Reviews     []Reviews     `gorm:"foreignKey:MoviesID"`
	// Reviews   []Reviews      `gorm:"foreignKey:MoviesID" json:"reviews"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
