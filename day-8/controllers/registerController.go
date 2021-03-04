package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type getBodyRegist struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Register func
func Register(c *gin.Context) {
	var body getBodyRegist
	err := c.Bind(&body)
	if err != nil {
		fmt.Println("Can't get data!")
	}

	// response
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "data successfully added.",
		"data": map[string]interface{}{
			"full_name": body.FullName,
			"email":     body.Email,
			"password":  body.Password,
		},
	})
}
