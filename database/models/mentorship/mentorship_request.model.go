package models

import (
	"log"
	"time"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MentorshipRequest struct {
	From     string    `bson:"from" json:"from"`
	To       string    `bson:"to" json:"to"`
	Date     time.Time `bson:"date" json:"date"`
	Time     time.Time `bson:"time" json:"time"`
	Duration int       `bson:"duration" json:"duration"`
	Note     string    `bson:"note" json:"note"`
}

func CreateMentorShipRequestCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"from", "to", "date", "time", "duration"},
		"properties": bson.M{
			"from": bson.M{
				"bsonType": "string",
			},
			"to": bson.M{
				"bsonType": "string",
			},
			"date": bson.M{
				"bsonType": "date",
			},
			"time": bson.M{
				"bsonType": "date",
			},
			"duration": bson.M{
				"bsonType": "int",
			},
			"note": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.MentorShipRequestCollection, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating MentorShip Request Collection: ", err)
		return
	}
	log.Printf("Mentorship Requests Collection Exists/Created Successfully")
}
