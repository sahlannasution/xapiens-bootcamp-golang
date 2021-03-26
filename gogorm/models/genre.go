package models

import "gorm.io/gorm"

// Genres struct
type Genres struct {
	Name        string        `json:"name"`
	MovieGenres []MovieGenres `gorm:"foreignKey:GenreID" json:"movie_genre"`
	gorm.Model
}
