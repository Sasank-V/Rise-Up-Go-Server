package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Review struct {
	From   string `bson:"from" json:"from"`
	To     string `bson:"to" json:"to"`
	Rating int    `bson:"rating" json:"rating"`
	Body   string `bson:"body" json:"body"`
}

func CreateReviewCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"from", "to", "rating", "body"},
		"properties": bson.M{
			"from": bson.M{
				"bsonType": "string",
			},
			"to": bson.M{
				"bsonType": "string",
			},
			"rating": bson.M{
				"bsonType": "int",
				"minimum":  1,
				"maximum":  5,
			},
			"body": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.ReviewCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating Review Collection: ", err)
		return
	}
	log.Printf("Review Collection Exists/Created Successfully")
}
