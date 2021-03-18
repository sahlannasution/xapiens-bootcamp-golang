package models

import "gorm.io/gorm"

// MovieGenres struct
type MovieGenres struct {
	MoviesID uint     `gorm:"primarykey" json:"movie_id"`
	Movies   []Movies `gorm:"foreignKey:ID" json:"movie"`
	GenresID uint     `gorm:"primarykey" json:"genre_id"`
	Genres   []Genres `gorm:"foreignKey:ID" json:"genre"`
	gorm.Model
}
