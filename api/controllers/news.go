package controllers

// Imports
import (
	"fmt"
	"log"
	"time"

	"../models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// NewsController represents the controller for operating on the News resource
	NewsController struct {
		session *mgo.Session
	}
)

// Database and Collection names
const (
	DB_NAME       = "hdgBball"
	DB_COLLECTION = "news"
)

// NewNewsController provides a reference to a NewsController with provided mongo session
func NewNewsController(s *mgo.Session) *NewsController {
	return &NewsController{s}
}

// checkErr when called: if error occurs, log to console
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// messageTypeDefault returns a successful message response code and body
func messageTypeDefault(msg string, c *gin.Context) {
	content := gin.H{
		"status": "200",
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

// checkErrTypeOne returns an error message and panic quits the API
func checkErrTypeOne(err error, msg string, status string, c *gin.Context) {
	if err != nil {
		panic(err)
		log.Fatalln(msg, err)
		content := gin.H{
			"status": status,
			"result": msg,
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, content)
	}
}

// NewsList returns every news object stored in Mongo
func (uc NewsController) NewsList(c *gin.Context) {
	var results []models.News

	// Query mongo with Find().All
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Find(nil).All(&results)
	if err != nil {
		checkErrTypeOne(err, "News doesn't exist", "403", c)
		return
	}

	// Reverses order of results, so that more recent news articles
	// appear first on the page.
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	c.JSON(200, results)
}

// CreateNews handles the HTTP request and calls child function to insert
// request (if valid) into the database
func (uc NewsController) CreateNews(c *gin.Context) {
	var json models.News
	c.Bind(&json)

	// Return results from createNews in the Response
	u := uc.createNews(json.Title, json.Body, c)
	if u.Title == json.Title {
		content := gin.H{
			"result": "Success",
			"Title":  u.Title,
			"Body":   u.Body,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

// DeleteNews searches mongo for a News entry to delete based off of the
// news entry's title
func (uc NewsController) DeleteNews(c *gin.Context) {
	var json models.News
	c.Bind(&json)

	// Call .Remove(title) in Mongo
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Remove(bson.M{"title": json.Title})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
	}

	content := gin.H{
		"result": "Success",
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

// createNews inserts the news object into the database
func (uc NewsController) createNews(Title string, Body string, c *gin.Context) models.News {

	// Declare and Format date the news entry was created
	currentTime := time.Now()
	date := currentTime.Format("01-02-2006")
	news := models.News{
		Title: Title,
		Body:  Body,
		Date:  date,
	}
	// Write the news entry to mongo
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Insert(&news)
	checkErrTypeOne(err, "Insert failed", "403", c)
	return news
}
