package main

import (
	// "xapiens-bootcamp-golang-day-11/config"
	// "xapiens-bootcamp-golang-day-11/controllers"
	// "xapiens-bootcamp-golang-day-11/models"

	"fmt"
	"xapiens-bootcamp-golang/day-12/config"
	"xapiens-bootcamp-golang/day-12/controllers"
	auth "xapiens-bootcamp-golang/day-12/middlewares/auth"
	"xapiens-bootcamp-golang/day-12/models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Sentry()
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// routing end point
	routing := gin.Default()

	// routing.POST("/login", authMiddlewares.LoginHandler)
	routing.POST("/login", auth.MiddleWare().LoginHandler)

	routing.NoRoute(auth.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		fmt.Println(jwt.ExtractClaims(c))
		// claims := jwt.ExtractClaims(c)

		// log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	requests := routing.Group("/requests")

	requests.Use(auth.MiddleWare().MiddlewareFunc())
	{

		requests.GET("/post", strDB.Post)
		requests.GET("/comment", strDB.Comment)
		requests.GET("/todos", strDB.Todos)

	}

	routing.Run()
}

// func Sentry() {
// 	var dsnSentry string

// 	if err := godotenv.Load(".env"); err != nil {
// 		// fmt.Println("This is the Error : ", err)
// 		sentry.CaptureException(err)
// 		log.Fatalf(err.Error())
// 	} else {
// 		dsnSentry = os.Getenv("DSN_SENTRY")
// 	}
// 	err := sentry.Init(sentry.ClientOptions{
// 		Dsn: dsnSentry,
// 	})
// 	if err != nil {
// 		log.Fatalf("sentry.Init: %s", err)
// 	}
// 	defer sentry.Flush(2 * time.Second)

// 	sentry.CaptureMessage("It works!")
// }
