package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type getBodyUpdate struct {
	FullName string `json:"fullName"`
}

//UpdateUserData func
func UpdateUserData(c *gin.Context) {
	var body getBodyRegist
	email := c.Query("email")
	errEmail := c.Bind(&email)
	err := c.Bind(&body)
	if err != nil || errEmail != nil {
		fmt.Println("User invalid!")
	}

	// response
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "data successfully updated.",
		"data": map[string]interface{}{
			"full_name": body.FullName,
			"email":     email,
			"role":      "user",
		},
	})
}
