package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type getBodyUserData struct {
	FullName string `json:"fullName"`
}

//GetUserData func
func GetUserData(c *gin.Context) {
	var body getBodyUserData
	err := c.Bind(&body)
	if err != nil {
		fmt.Println("Can't get data!")
	}

	// response
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "Successfully Get Data.",
		"data": map[string]interface{}{
			"full_name": body.FullName,
		},
	})
}
