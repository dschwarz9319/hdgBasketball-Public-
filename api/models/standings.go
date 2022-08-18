package models

// Imports
import "gopkg.in/mgo.v2/bson"

type (
	// Data structure for a Team as well as their standings within a league
	Standings struct {
		DBId        bson.ObjectId `json:"DbId" bson:"_id,omitempty"`
		Team        string        `json:"team" bson:"team"`     // Required
		League      string        `json:"league" bson:"league"` // Required
		Wins        int           `json:"wins" bson:"wins,omitempty"`
		Losses      int           `json:"losses" bson:"losses,omitempty"`
		GamesPlayed int           `json:"gamesPlayed" bson:"gamesPlayed,omitempty"`
		Ratio       float64       `json:"ratio" bson:"ratio,omitempty"`
	}
)
