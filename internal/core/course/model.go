package course

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Beginner     Difficulty = "beginner"
	Intermediate Difficulty = "intermediate"
	Advanced     Difficulty = "advanced"
)

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

const (
	Video    ContentType = "video"
	Audio    ContentType = "audio"
	Document ContentType = "document"
)

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
			"order_no": bson.M{
				"bsonType": "int",
			},
		},
	}
	err := database.CreateCollection(db, lib.LessonCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Lessons Collection: ", err)
		return
	}
	log.Println("Lessons Collection Created Successfully")
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
			"order_no": bson.M{
				"bsonType": "int",
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
	log.Println("Module Collection created Successfully")
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
	log.Println("Resource Collection Created Successfully")
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
	log.Println("Course Progress Collection Exists/Created Successfully")
}

func ConnectAllCourseCollections() {
	ConnectCourseCollection()
	ConnectModuleCollection()
	ConnectLessonCollection()
	ConnectResourceCollection()
	ConnectCourseProgressCollection()
}
