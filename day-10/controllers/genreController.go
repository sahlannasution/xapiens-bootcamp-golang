package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Genres struct
type Genres struct {
	ID   string `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

// AddGenre func
func (StrDB *StrDB) AddGenre(c *gin.Context) {
	var (
		genre  Genres
		result gin.H
	)
	name := c.PostForm("name")

	genre.Name = name
	StrDB.DB.Create(&genre)
	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Register!",
		"data":    genre,
	}

	c.JSON(http.StatusOK, result)
}

// GetGenresList func
func (StrDB *StrDB) GetGenresList(c *gin.Context) {
	var (
		genres []Genres
		result gin.H
	)

	StrDB.DB.Find(&genres)
	result = gin.H{
		"status":  "success",
		"message": "Successfully Get Genre List",
		"data":    genres,
	}

	c.JSON(http.StatusOK, result)
}
