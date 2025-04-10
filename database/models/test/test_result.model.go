package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TestResult struct {
	TestID   string `bson:"test_id" json:"test_id"`
	UserID   string `bson:"user_id" json:"user_id"`
	Result   int    `bson:"result" json:"result"`
	FeedBack string `bson:"feedback" json:"feedback"`
}

func CreateTestResultCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"test_id", "user_id", "result"},
		"properties": bson.M{
			"test_id": bson.M{
				"bsonType": "string",
			},
			"user_id": bson.M{
				"bsonType": "string",
			},
			"result": bson.M{
				"bsonType": "int",
			},
			"feedback": bson.M{
				"bsonType": "string",
			},
		},
	}

	err := database.CreateCollection(db, lib.TestResultCollectionName, jsonSchema, []string{"test_id", "user_id"})
	if err != nil {
		log.Fatal("Error creating TestResult Collection: ", err)
		return
	}
	log.Printf("TestResult Collection Exists/Created Successfully\n")
}
