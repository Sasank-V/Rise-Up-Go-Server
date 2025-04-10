package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionType string

const (
	Private SessionType = "private"
	Public  SessionType = "public"
)

type MentorShipSession struct {
	MentorID  string   `bson:"mentor_id" json:"mentor_id"`
	RequestID string   `bson:"request_id" json:"request_id"`
	Link      string   `bson:"session_link" json:"session_link"`
	Title     string   `bson:"title" json:"title"`
	Resources []string `bson:"resources" json:"resources"`
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
	log.Printf("Mentorship Session Collection Exists/Created Successfully\n")
}
