package controllers

import (
	"net/http"
	logger "xapiens-bootcamp-golang/day-13/log"
	"xapiens-bootcamp-golang/day-13/models"

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
		genre  models.Genres
		result gin.H
	)
	name := c.PostForm("name")

	if name == "" {
		e := "Field is required!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		// fmt.Println("Field Email, Password, FullName, Role is required!")
		c.JSON(http.StatusBadRequest, result)

		// logger.Sentry(err) // push log error ke sentry
	} else {
		genre.Name = name
		if res := StrDB.DB.Create(&genre); res.Error != nil {
			err := res.Error
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "bad request",
				"message": err.Error(),
			})

			logger.Sentry(err) // push log error ke sentry
		} else {
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Register!",
				"data":    genre,
			}

			c.JSON(http.StatusOK, result)
		}
	}
}

// GetGenresList func
func (StrDB *StrDB) GetGenresList(c *gin.Context) {
	var (
		genres []Genres
		result gin.H
	)

	if res := StrDB.DB.Find(&genres); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "not found",
			"message": err.Error(),
		}

		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err) // push log error ke sentry
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Successfully Get Genre List",
			"data":    genres,
		}

		c.JSON(http.StatusOK, result)
	}

}
