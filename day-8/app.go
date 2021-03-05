package main

import (
	"log"
	"xapiens-bootcamp-golang/day-8/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	//ini ngikut dari dokumentasi gin
	// request.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	request.POST("/register", controllers.Register)
	request.POST("/login", controllers.Login)
	request.GET("/user/", controllers.GetUserData)
	request.PUT("/user/", controllers.UpdateUserData)
	request.GET("/movie/", controllers.GetDetailMovie)

	log.Println("Server up and run on Port 8080")
	request.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"))
}
