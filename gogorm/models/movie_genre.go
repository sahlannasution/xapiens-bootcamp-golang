package models

import "gorm.io/gorm"

// MovieGenres struct
type MovieGenres struct {
	MovieID uint `json:"movie_id"`

	GenreID uint ` json:"genre_id"`
	gorm.Model
}
