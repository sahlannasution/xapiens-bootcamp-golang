package models

import "gorm.io/gorm"

// Reviews struct
type Reviews struct {
	ID      uint   `gorm:"primarykey, autoIncrement" json:"id"`
	Review  string `json:"review"`
	Rate    int    `json:"rate"`
	UserID  int    `json:"user_id"`
	Users   Users  `gorm:"foreignKey:ID" json:"user"`
	MovieID int    `json:"movie_id"`
	Movies  Movies `gorm:"foreignKey:ID" json:"movie"`
	gorm.Model
}
