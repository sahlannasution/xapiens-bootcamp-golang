package models

import "gorm.io/gorm"

// MovieGenre struct
type MovieGenre struct {
	GenreID string `json:"genre_id"`
	MovieID string `json:"movie_id"`
	// Genre   []Genres `gorm:"ForeignKey:GenreID"`
	// Movies  []Movies `gorm:"ForeignKey:MovieID"`
	gorm.Model
}
