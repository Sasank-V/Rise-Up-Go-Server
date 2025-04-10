package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CourseProgress struct {
	UserID           string   `bson:"user_id" json:"user_id"`
	CourseID         string   `bson:"course_id" json:"course_id"`
	LessonsCompleted []string `bson:"lessons_completed" json:"lessons_completed"`
	CourseCompleted  bool     `bson:"course_completed" json:"course_completed"`
}

func CreateCourseProgressCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "course_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"course_id": bson.M{
				"bsonType": "string",
			},
			"lessons_completed": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"course_completed": bson.M{
				"bsonType": "bool",
			},
		},
	}
	err := database.CreateCollection(db, lib.CourseProgressCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating Course Progress Collection")
		return
	}
	log.Printf("Course Progress Collection Exists/Created Successfully")
}
