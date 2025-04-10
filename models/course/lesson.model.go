package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ContentType string

const (
	Video    ContentType = "video"
	Audio    ContentType = "audio"
	Document ContentType = "document"
)

type Lesson struct {
	ModuleID    string      `bson:"module_id" json:"module_id"`
	Title       string      `bson:"title" json:"title"`
	Description string      `bson:"description" json:"description"`
	ContentLink string      `bson:"content_link" json:"content_link"`
	ContentType ContentType `bson:"content_type" json:"content_type"`
	Resources   []string    `bson:"resources" json:"resources"`
	Duration    int         `bson:"duration" json:"duration"`
}

func CreateLessonCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"module_id", "title", "description"},
		"properties": bson.M{
			"module_id": bson.M{
				"bsonType": "string",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"description": bson.M{
				"bsonType": "string",
			},
			"content_link": bson.M{
				"bsonType": "string",
			},
			"content_type": bson.M{
				"bsonType": "string",
				"enum":     []string{string(Video), string(Audio), string(Document)},
			},
			"resources": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"duration": bson.M{
				"bsonType": "int",
			},
		},
	}
	err := database.CreateCollection(db, lib.LessonCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Lessons Collection: ", err)
		return
	}
	log.Printf("Lessons Collection Created Successfully")
}
