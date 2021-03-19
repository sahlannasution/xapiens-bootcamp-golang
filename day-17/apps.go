package main

import (
	"log"
	"xapiens-bootcamp-golang/day-17/config"
	"xapiens-bootcamp-golang/day-17/controllers"
	"xapiens-bootcamp-golang/day-17/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

// type RegistMailer struct {
// 	ID          uint      `gorm:"primarykey, autoIncrement" json:"id"`
// 	Type        string    `json:"type"`
// 	Email       string    `json:"email"`
// 	Message     string    `json:"message"`
// 	Status      bool      `json:"status"`
// 	DeliveredAt time.Time `json:"delivered_at"`
// }
// StrDB struct
// type StrinDB struct {
// 	DataB *gorm.DB
// }

func main() {
	// Db Connections
	dbPG := config.Connection()
	StrDB := controllers.StrDB{DB: dbPG}
	// StrinDB := StrinDB{DataB: dbPG}

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
	routing.POST("/register", StrDB.Register)

	/* User Signin */
	routing.POST("/signin", StrDB.MiddleWare().LoginHandler)

	routing.NoRoute(StrDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	movie_reviews := routing.Group("/movie_reviews")

	movie_reviews.Use(StrDB.MiddleWare().MiddlewareFunc())
	{
		/* User Routes */
		movie_reviews.GET("/user/:email", StrDB.GetUserData)
		movie_reviews.PUT("/user/:email", StrDB.UpdateUser)
		/* End User Routes */

		/* Movies Routes */
		movie_reviews.POST("/movies", StrDB.AddMovies)
		movie_reviews.GET("/movies", StrDB.GetMoviesList)
		/* End Movies Routes */

		/* Genre Routes */
		movie_reviews.POST("/genre", StrDB.AddGenre)
		movie_reviews.GET("/genre", StrDB.GetGenresList)
		/* End Genre Routes */

		/* Movies Genre Routes */
		movie_reviews.POST("/movies/genre", StrDB.AddMoviesGenre)
		/* End Movies Genre Routes */

		/* Review Routes */
		movie_reviews.POST("/review", StrDB.AddReview)
		movie_reviews.GET("/review/:movie_id", StrDB.GetReviewByMovie)
		movie_reviews.GET("/reviews/detail/:review_id", StrDB.GetReviewDetail)
		/* End Review Routes */
	}

	go func() {

		gocron.Every(1).Minute().Do(StrDB.Job)
		<-gocron.Start()
	}()
	routing.Run(":6969")

}

// func cronJob() {

// }
