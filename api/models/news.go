package models

// Imports
import "gopkg.in/mgo.v2/bson"

type (
	// Data structures for a News post
	News struct {
		DBId  bson.ObjectId `json:"DbId" bson:"_id,omitempty"`
		Title string        `json:"title" bson:"title"` // Required
		Body  string        `json:"body" bson:"body"`   // Required
		Date  string        `json:"date" bson:"date,omitempty"`
	}
)
