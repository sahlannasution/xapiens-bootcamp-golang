package models

import "gorm.io/gorm"

// Reviews struct
type Reviews struct {
	ID       uint   `gorm:"primarykey, autoIncrement" json:"id"`
	Review   string `json:"review"`
	Rate     int    `json:"rate"`
	UsersID  uint   `json:"users_id"`
	Users    Users  `gorm:"foreignKey:ID" json:"users"`
	MoviesID uint   `json:"movies_id"`
	Movies   Movies `gorm:"foreignKey:ID" json:"movies"`
	gorm.Model
}
