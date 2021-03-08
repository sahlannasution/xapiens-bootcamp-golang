package models

import "gorm.io/gorm"

// Movies struct
type Movies struct {
	Title      string        `json:"title"`
	Year       int           `json:"year"`
	Ratings    int           `json:"ratings"`
	MovieGenre []MovieGenres `gorm:"foreignKey:MoviesID" json:"movie_genres"`
	gorm.Model
}
