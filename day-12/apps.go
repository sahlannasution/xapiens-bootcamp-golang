package main

import (
	// "xapiens-bootcamp-golang-day-11/config"
	// "xapiens-bootcamp-golang-day-11/controllers"
	// "xapiens-bootcamp-golang-day-11/models"

	"log"
	"xapiens-bootcamp-golang/day-12/config"
	"xapiens-bootcamp-golang/day-12/controllers"
	middlewares "xapiens-bootcamp-golang/day-12/middlewares/auth"
	"xapiens-bootcamp-golang/day-12/models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// routing end point
	routing := gin.Default()

	// routing.POST("/login", authMiddlewares.LoginHandler)
	routing.POST("/login", middlewares.MiddleWare().LoginHandler)

	routing.NoRoute(middlewares.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	requests := routing.Group("/requests")

	requests.Use(middlewares.MiddleWare().MiddlewareFunc())
	{

		requests.GET("/post", strDB.Post)
		requests.GET("/comment", strDB.Comment)
		requests.GET("/todos", strDB.Todos)
		// auth.GET("/hello", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"text": "Berhasil masuk ke JWT",
		// 	})
		// })

		// auth.GET("/vendorList", strDB.GetVendorList)

		// // routing vendor
		// auth.POST("/vendor", strDB.PostDataVendor)

		// auth.GET("/VendorByQuery", strDB.GetDataVendor)

		// // routing employee
		// auth.POST("/employee", strDB.PostDataEmployee)
		// auth.GET("/employeeList", strDB.GetEmployeeList)

		// // routing upload files
		// auth.POST("/uploadFile", controller.UploadSingleFile)
		// auth.POST("/uploadMultipleFile", controller.UploadMultipleFile)

		// //routing 3rd api / http request
		// auth.GET("/post", controller.HttpRequest)
	}

	//routing 3rd api / http request

	routing.Run()
}
