package models

import "gorm.io/gorm"

// Genres struct
type Genres struct {
	Name string `json:"name"`
	// Movies []Movies `gorm:"many2many:movie_genres"`
	gorm.Model
}
