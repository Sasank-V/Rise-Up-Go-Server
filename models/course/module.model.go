package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Module struct {
	CourseID string   `bson:"course_id" json:"title_id"`
	Title    string   `bson:"title" json:"title"`
	Lessons  []string `bson:"lessons" json:"lessons"`
}

func CreateModuleCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"course_id", "title"},
		"properties": bson.M{
			"course_id": bson.M{
				"bsonType": "string",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"lessons": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}
	err := database.CreateCollection(db, lib.ModuleCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Module Collection: ", err)
		return
	}
	log.Printf("Moduel Collection created Successfully")
}
