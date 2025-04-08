package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Difficulty string

const (
	Beginner     Difficulty = "beginner"
	Intermediate Difficulty = "intermediate"
	Advanced     Difficulty = "advanced"
)

type Course struct {
	Owner       string     `bson:"owner" json:"owner"`
	Banner      string     `bson:"banner" json:"banner"`
	Title       string     `bson:"title" json:"title"`
	Description string     `bson:"description" json:"description"`
	Difficulty  Difficulty `bson:"difficulty" json:"difficulty"`
	Duration    int        `bson:"duration" json:"duration"`
	Modules     []string   `bson:"modules" json:"modules"`
	Instructors []string   `bson:"instructors" json:"instructors"`
	Discussions []string   `bson:"discussion" json:"discussions"`
}

func CreateCourseCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{},
		"properties": bson.M{
			"owner": bson.M{
				"bsonType": "string",
			},
			"banner": bson.M{
				"bsonType": "string",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"description": bson.M{
				"bsonType": "string",
			},
			"difficulty": bson.M{
				"bsonType": "string",
			},
			"duration": bson.M{
				"bsonType": "int",
			},
			"modules": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"instructors": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"discussions": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}
	err := database.CreateCollection(db, lib.CourseCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating Course Collection: ", err)
		return
	}
	log.Printf("Course Collection Created Successfully")
}
