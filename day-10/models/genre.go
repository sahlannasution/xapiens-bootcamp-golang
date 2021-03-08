package models

import "gorm.io/gorm"

// Genres struct
type Genres struct {
	Name string `json:"name"`
	gorm.Model
}
