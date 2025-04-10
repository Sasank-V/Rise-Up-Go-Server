package test

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TestType string

const (
	MockInterview TestType = "mockinterview"
	SkillQuiz     TestType = "skillquiz"
)

type Test struct {
	Type       TestType `bson:"type" json:"type"`
	CourseID   string   `bson:"course_id" json:"course_id"`
	Skills     []string `bson:"skills" json:"skills"`
	Difficulty string   `bson:"difficulty" json:"difficulty"`
	Questions  []string `bson:"questions" json:"questions"`
}

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
	log.Printf("Test Collection Exists/Created Successfully\n")
}
