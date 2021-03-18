package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	logger "xapiens-bootcamp-golang/day-17/log"
	"xapiens-bootcamp-golang/day-17/models"

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
		e := "Field name is required!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		// fmt.Println("Field Email, Password, FullName, Role is required!")
		c.JSON(http.StatusBadRequest, result)

		logger.SentryStr(e) // push log error ke sentry
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

	if check, data := CheckRedis("genres"); check {
		if err := json.Unmarshal(data, &genres); err != nil {
			fmt.Println("Error", err.Error())
		}
		fmt.Println(genres)
		result = gin.H{
			"status":  "success",
			"message": "Successfully Get Genre List",
			"data":    genres,
		}
		c.JSON(http.StatusOK, result)

	} else {
		if res := StrDB.DB.Find(&genres); res.Error != nil {
			err := res.Error
			result = gin.H{
				"status":  "not found",
				"message": err.Error(),
			}

			c.JSON(http.StatusNotFound, result)
			logger.Sentry(err) // push log error ke sentry
		} else {
			output, _ := json.Marshal(genres)
			result = gin.H{
				"status":  "success",
				"message": "Successfully Get Genre List",
				"data":    genres,
			}

			c.JSON(http.StatusOK, result)
			SetRedis("genres", string(output))
		}
	}

}
