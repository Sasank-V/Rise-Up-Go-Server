package test

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MockInterview TestType = "mockinterview"
	SkillQuiz     TestType = "skillquiz"
)

func CreateTestCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"type", "course_id", "skills", "questions"},
		"properties": bson.M{
			"type": bson.M{
				"bsonType": "string",
				"enum":     []string{string(MockInterview), string(SkillQuiz)},
			},
			"course_id": bson.M{
				"bsonType": "string",
			},
			"skills": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"diffuculty": bson.M{
				"bsonType": "string",
			},
			"questions": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}

	err := database.CreateCollection(db, lib.TestCollectionName, jsonSchema, []string{"course_id", "type"})
	if err != nil {
		log.Fatal("Error creating Test Collection: ", err)
		return
	}
	log.Println("Test Collection Exists/Created Successfully")
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
	log.Println("TestResult Collection Exists/Created Successfully")
}

func CreateAllTestCollections(db *mongo.Database) {
	CreateTestCollection(db)
	CreateTestResultCollection(db)
}
