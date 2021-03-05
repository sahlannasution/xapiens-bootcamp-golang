package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetDetailMovie func
func GetDetailMovie(c *gin.Context) {
	// var body getBodyUserData
	id := c.Query("id")
	err := c.Bind(&id)
	if err != nil {
		fmt.Println("Wrong Id")
	}

	// response
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Successfully Get Detail Movie",
		"data": map[string]interface{}{
			"id":      id,
			"title":   "Parasyte",
			"year":    "2019",
			"ratings": 8,
		},
	})
}
