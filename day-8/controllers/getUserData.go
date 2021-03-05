package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetUserData func
func GetUserData(c *gin.Context) {
	// var body getBodyUserData
	email := c.Query("email")
	err := c.Bind(&email)
	if err != nil {
		fmt.Println("Wrong Email")
	}

	// response
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "Successfully Get Data.",
		"data": map[string]interface{}{
			"full_name": "Sahlan Nasution",
			"email":     email,
			"role":      "user",
		},
	})
}
