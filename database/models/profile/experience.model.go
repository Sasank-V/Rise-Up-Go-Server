package models

import (
	"log"
	"time"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Experience struct {
	UserID      string    `bson:"user_id" json:"user_id"`
	Company     string    `bson:"company" json:"company"`
	Position    string    `bson:"position" json:"position"`
	StartDate   time.Time `bson:"start_date" json:"start_date"`
	EndDate     time.Time `bson:"end_date" json:"end_date"`
	Description string    `bson:"description" json:"description"`
}

func CreateExperienceCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "company", "position", "start_date"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"company": bson.M{
				"bsonType": "string",
			},
			"position": bson.M{
				"bsonType": "string",
			},
			"start_date": bson.M{
				"bsonType": "date",
			},
			"end_date": bson.M{
				"bsonType": "date",
			},
			"description": bson.M{
				"bsonType": "string",
			},
		},
	}

	err := database.CreateCollection(db, lib.ExperienceCollectionName, jsonSchema, []string{"user_id", "company", "position"})
	if err != nil {
		log.Fatal("Error creating Experience Collection: ", err)
		return
	}
	log.Printf("Experience Collection Exists/Created Successfully\n")
}
