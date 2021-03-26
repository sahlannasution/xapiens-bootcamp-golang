package models

import "gorm.io/gorm"

// Movies struct
type Movies struct {
	Title       string        `json:"title"`
	Year        int           `json:"year"`
	Ratings     int           `json:"ratings"`
	MovieGenres []MovieGenres `gorm:"foreignKey:MovieID" json:"movie_genre"`
	// Reviews []Reviews
	gorm.Model
}
