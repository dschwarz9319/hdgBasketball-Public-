package controllers

// Imports
import (
	"fmt"
	"math"

	"../models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// StandingsController represents the controller for operating on the Standings resource
	StandingsController struct {
		session *mgo.Session
	}
)

// Database and Collection names

const (
	STANDINGS_DB_NAME       = "hdgBball"
	STANDINGS_DB_COLLECTION = "standings"
)

// NewStandingsController provides a reference to a StandingsController with provided mongo session
func NewStandingsController(s *mgo.Session) *StandingsController {
	return &StandingsController{s}
}

// StandingsList returns every standings object stored in Mongo
func (uc StandingsController) StandingsList(c *gin.Context) {
	var results []models.Standings

	// Query mongo with Find().All
	err := uc.session.DB(STANDINGS_DB_NAME).C(STANDINGS_DB_COLLECTION).Find(nil).All(&results)
	if err != nil {
		checkErrTypeOne(err, "Standings doesn't exist", "403", c)
		return
	}
	c.JSON(200, results)
}

// CreateStandings creates a new standings object and calls a child function
// to insert it into the database
func (uc StandingsController) CreateStandings(c *gin.Context) {
	var json models.Standings
	c.Bind(&json)

	// Return results from createStandings in the response
	u := uc.createStandings(json.Team, json.League, json.Wins, json.Losses, c)
	if u.Team == json.Team {
		content := gin.H{
			"result": "Success",
			"Team":   u.Team,
			"League": u.League,
			"Wins":   u.Wins,
			"Losses": u.Losses,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

// UpdateStandings updates an existing standings entry in the database
func (uc StandingsController) UpdateStandings(c *gin.Context) {
	var json models.Standings
	var ratio float64
	var ratioFlag bool
	ratioFlag = false
	c.Bind(&json)

	if json.Wins == 0 {
		ratio = 0
		ratioFlag = true
	}

	if json.Losses == 0 {
		ratio = 1
		ratioFlag = true
	}

	// Calculate win loss ratio
	if ratioFlag == false {
		ratio = 1 - (float64(json.Losses) / float64(json.Wins+json.Losses))
		ratio = math.Round(ratio*100) / 100
	}

	standings := models.Standings{
		Team:        json.Team,
		League:      json.League,
		Wins:        json.Wins,
		Losses:      json.Losses,
		GamesPlayed: json.Wins + json.Losses,
		Ratio:       ratio,
	}

	// Update entry in the database by searching for the request's team name
	err := uc.session.DB(STANDINGS_DB_NAME).C(STANDINGS_DB_COLLECTION).Update(bson.M{"team": json.Team, "league": json.League}, &standings)
	if err != nil {
		fmt.Printf("Update fail %v\n", err)
	}

	content := gin.H{
		"result": "Success",
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)

}

// DeleteStandings deletes a standings entry in the database
func (uc StandingsController) DeleteStandings(c *gin.Context) {
	var json models.Standings
	c.Bind(&json)

	// Search for entry by the team's name
	err := uc.session.DB(STANDINGS_DB_NAME).C(STANDINGS_DB_COLLECTION).Remove(bson.M{"team": json.Team, "league": json.League})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
	}

	content := gin.H{
		"result": "Success",
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

// createStandings inserts a new standings object into the database
func (uc StandingsController) createStandings(Team string, League string, Wins int, Losses int, c *gin.Context) models.Standings {

	standings := models.Standings{
		Team:        Team,
		League:      League,
		Wins:        Wins,
		Losses:      Losses,
		GamesPlayed: Wins + Losses,
		Ratio:       0,
	}
	// Write the standings to mongo
	err := uc.session.DB(STANDINGS_DB_NAME).C(STANDINGS_DB_COLLECTION).Insert(&standings)
	checkErrTypeOne(err, "Insert failed", "403", c)
	return standings
}
