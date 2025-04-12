package mentorship

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Private SessionType = "private"
	Public  SessionType = "public"
)

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
			"type": bson.M{
				"bsonType": "string",
				"enum":     []string{string(Private), string(Public)},
			},
			"note": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.MentorShipRequestCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating MentorShip Request Collection: ", err)
		return
	}
	log.Println("Mentorship Requests Collection Exists/Created Successfully")
}

func CreateMentorShipSessionCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"mentor_id", "request_id", "session_link", "title"},
		"properties": bson.M{
			"mentor_id": bson.M{
				"bsonType": "string",
			},
			"request_id": bson.M{
				"bsonType": "string",
			},
			"session_link": bson.M{
				"bsonType": "string",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"resources": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}

	err := database.CreateCollection(db, lib.MentorShipSessionCollectionName, jsonSchema, []string{"mentor_id", "request_id"})
	if err != nil {
		log.Fatal("Error creating Mentorship Session Collection: ", err)
		return
	}
	log.Println("Mentorship Session Collection Exists/Created Successfully")
}

func CreateAllMentorshipCollections(db *mongo.Database) {
	CreateMentorShipRequestCollection(db)
	CreateMentorShipSessionCollection(db)
}
