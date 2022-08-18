package main

// Imports
import (
	"./controllers"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gopkg.in/mgo.v2"
)

// main is the endpoint router of the API. Upon receiving an HTTP request along with
// an endpoint, main calls a function with the proper controller resource
func main() {

	// Declaration of controllers for different endpoints and their data structures
	newsController := controllers.NewNewsController(getSession())
	standingsController := controllers.NewStandingsController(getSession())

	// Allow for access to API using CORS
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*", // TODO: in production, change to only allow access from localhost
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Set-Cookie",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// News endpoint CRUD operations
	router.GET("/news", newsController.NewsList)
	router.POST("/news", newsController.CreateNews)
	router.DELETE("/news", newsController.DeleteNews)

	// Standings endpoint CRUD operations
	router.GET("/standings", standingsController.StandingsList)
	router.POST("/standings", standingsController.CreateStandings)
	router.PUT("/standings", standingsController.UpdateStandings)
	router.DELETE("/standings", standingsController.DeleteStandings)

	// TODO: add endpoints for 'schedule' HTTP requests

	// Run router on port 8000
	router.Run(":8000")
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {

	mode := gin.Mode()

	// Prod Mongo address
	if mode == "release" {
		s, err := mgo.Dial("mongodb://mongo")
		if err != nil {
			panic(err)
		}
		return s
	}

	// Dev Mongo address
	if mode == "debug" {
		s, err := mgo.Dial("mongodb://localhost")
		if err != nil {
			panic(err)
		}
		return s
	}

	// Default case if ENV mode variable is no set
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
