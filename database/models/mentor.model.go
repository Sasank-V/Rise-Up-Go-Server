package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Mentor struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	Skills             []string `bson:"skills" json:"skills"`
	Experience         []string `bson:"experience" json:"experience"`
	RegisteredCourses  []string `bson:"registered_courses" json:"registered_courses"`
	MentorshipRequests []string `bson:"mentorship_requests" json:"mentorship_requests"`
	MentorShipSessions []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
	TestsTaken         []string `bson:"tests_taken" json:"tests_taken"`
	Reviews            []string `bson:"reviews" json:"reviews"`
	Available          bool     `bson:"available" json:"available"`
}

func CreateMentorCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"registered_courses": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"mentorship_requests": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"mentoship_sessions": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"tests_taken": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"reviews": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"available": bson.M{
				"bsonType": "bool",
			},
		},
	}
	err := database.CreateCollection(db, lib.MentorCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error Creating the Mentor Collection: ", err)
		return
	}
	log.Printf("Mentor Collection Exists/Created Successfully")
}
