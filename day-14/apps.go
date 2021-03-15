package main

import (
	"log"
	"xapiens-bootcamp-golang/day-14/config"
	"xapiens-bootcamp-golang/day-14/controllers"
	"xapiens-bootcamp-golang/day-14/models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Db Connections
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	// Migrate models to DB Postgre
	models.Migrations(dbPG)

	// Seeding Data
	models.SeederUser(dbPG)
	models.SeederGenre(dbPG)
	models.SeederMovie(dbPG)

	routing := gin.Default()

	/* User Register */
	routing.POST("/register", strDB.Register)

	/* User Signin */
	routing.POST("/signin", strDB.MiddleWare().LoginHandler)

	routing.NoRoute(strDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	movie_reviews := routing.Group("/movie_reviews")

	movie_reviews.Use(strDB.MiddleWare().MiddlewareFunc())
	{
		/* User Routes */
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
		/* End Movies Genre Routes */

		/* Review Routes */
		routing.POST("/review", strDB.AddReview)
		routing.GET("/review/:movie_id", strDB.GetReviewByMovie)
		routing.GET("/reviews/detail/:review_id", strDB.GetReviewDetail)
		/* End Review Routes */
	}

	routing.Run(":6969")
}
