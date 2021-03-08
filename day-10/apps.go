package main

import (

	// "xapiens-day9/controller"
	// "xapiens-day9/models"

	"xapiens-bootcamp-golang/day-10/config"
	"xapiens-bootcamp-golang/day-10/controllers"
	"xapiens-bootcamp-golang/day-10/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Db Connections
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	models.Migrations(dbPG)

	routing := gin.Default()

	/* User Routes */
	routing.POST("/register", strDB.Register)
	routing.POST("/signin", strDB.Signin)
	routing.GET("/user/:email", strDB.GetUserData)
	routing.PUT("/user/:email", strDB.UpdateUser)
	/* End User Routes */

	/* Movies Routes */
	routing.POST("/movies", strDB.AddMovies)
	routing.GET("/movies", strDB.GetMoviesList)
	/* End Movies Routes */

	/* Genre Routes */
	routing.POST("/genre", strDB.AddGenre)
	routing.GET("/genre", strDB.GetGenresList)
	/* End Genre Routes */

	/* Movies Genre Routes */
	routing.POST("/movies/genre", strDB.AddMoviesGenre)
	// routing.GET("/movies/genre", strDB.GetMoviesGenre)

	/* End Movies Genre Routes */

	routing.Run(":6969")
}
