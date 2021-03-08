package models

import "gorm.io/gorm"

// Genres struct
type Genres struct {
	Name       string        `json:"name"`
	MovieGenre []MovieGenres `gorm:"foreignKey:GenresID" json:"movie_genres"`
	gorm.Model
}
