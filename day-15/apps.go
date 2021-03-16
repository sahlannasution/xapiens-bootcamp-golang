package main

import (
	"log"
	"xapiens-bootcamp-golang/day-15/config"
	"xapiens-bootcamp-golang/day-15/controllers"
	"xapiens-bootcamp-golang/day-15/models"

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
	models.SeederMovieGenre(dbPG)
	models.SeederReview(dbPG)

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
		movie_reviews.GET("/user/:email", strDB.GetUserData)
		movie_reviews.PUT("/user/:email", strDB.UpdateUser)
		/* End User Routes */

		/* Movies Routes */
		movie_reviews.POST("/movies", strDB.AddMovies)
		movie_reviews.GET("/movies", strDB.GetMoviesList)
		/* End Movies Routes */

		/* Genre Routes */
		movie_reviews.POST("/genre", strDB.AddGenre)
		movie_reviews.GET("/genre", strDB.GetGenresList)
		/* End Genre Routes */

		/* Movies Genre Routes */
		movie_reviews.POST("/movies/genre", strDB.AddMoviesGenre)
		/* End Movies Genre Routes */

		/* Review Routes */
		movie_reviews.POST("/review", strDB.AddReview)
		movie_reviews.GET("/review/:movie_id", strDB.GetReviewByMovie)
		movie_reviews.GET("/reviews/detail/:review_id", strDB.GetReviewDetail)
		/* End Review Routes */
	}

	routing.Run(":6969")
}
