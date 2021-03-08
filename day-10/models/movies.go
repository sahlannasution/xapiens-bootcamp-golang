package models

import "gorm.io/gorm"

// Movies struct
type Movies struct {
	Title   string `json:"title"`
	Year    int    `json:"year"`
	Ratings int    `json:"ratings"`
	gorm.Model
}
