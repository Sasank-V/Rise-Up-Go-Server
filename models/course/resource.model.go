package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Resource struct {
	LessonID string `bson:"lesson_id" json:"lesson_id"`
	Name     string `bson:"name" json:"name"`
	Link     string `bson:"link" json:"link"`
}

func CreateResourceCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"lesson_id", "name", "link"},
		"properties": bson.M{
			"lesson_id": bson.M{
				"bsonType": "string",
			},
			"name": bson.M{
				"bsonType": "string",
			},
			"link": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.ResourceCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Resource Collection: ", err)
		return
	}
	log.Printf("Resource Collection Created Successfully")
}
