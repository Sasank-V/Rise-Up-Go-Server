package course

import (
	"log"
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var CourseColl *mongo.Collection
var courseConnect sync.Once

func ConnectCourseCollection() {
	courseConnect.Do(func() {
		db := database.InitDB()
		CreateCourseCollection(db)
		CourseColl = db.Collection(lib.CourseCollectionName)
	})
}

func CreateCourseCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"owner", "banner", "title", "difficulty", "description", "duration"},
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
			"skills": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
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
	log.Print("Course Collection Exists/Created Successfully")
}