package profile

import (
	"log"
	"time"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Education struct {
	UserID    string    `bson:"user_id" json:"user_id"`
	Institute string    `bson:"institute" json:"institute"`
	Degree    string    `bson:"degree" json:"degree"`
	StartDate time.Time `bson:"start_date" json:"start_date"`
	EndDate   time.Time `bson:"end_date" json:"end_date"`
}

func CreateEducationCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "institute", "degree", "start_date"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"institute": bson.M{
				"bsonType": "string",
			},
			"degree": bson.M{
				"bsonType": "string",
			},
			"start_date": bson.M{
				"bsonType": "date",
			},
			"end_date": bson.M{
				"bsonType": "date",
			},
		},
	}
	err := database.CreateCollection(db, lib.EducationCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Education Collection: ", err)
		return
	}
	log.Printf("Education Collection Exists/Created Successfully")
}
