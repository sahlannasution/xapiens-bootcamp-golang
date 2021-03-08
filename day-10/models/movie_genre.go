package models

import "gorm.io/gorm"

// MovieGenres struct
type MovieGenres struct {
	MoviesID uint     `gorm:"primarykey" json:"movie_id"`
	Movie    []Movies `gorm:"foreignKey:ID"`
	// MovieName string `json:"name"`
	// GenresID uint `json:"genres_id"`
	gorm.Model
}
