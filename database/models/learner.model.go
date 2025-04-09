package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Learner struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	Skills             string   `bson:"skills" json:"skills"`
	Interests          string   `bson:"interests" json:"interests"`
	JobPreferences     []string `bson:"job_prefrences" json:"job_preferences"`
	LanguagePreferred  string   `bson:"language_preferred" json:"language_preferred"`
	EnrolledCourses    []string `bson:"enrolled_courses" json:"enrolled_courses"`
	AppliedJobs        []string `bson:"applied_jobs" json:"applied_jobs"`
	TestsTaken         []string `bson:"tests_taken" json:"tests_taken"`
	MentorshipSessions []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
}

func CreateLearnerCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"enrolled_courses": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"applied_jobs": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"tests_taken": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"mentorship_sessions": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}
	err := database.CreateCollection(db, lib.LeanerCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error creating Learner Collection: ", err)
		return
	}
	log.Printf("Learner Collection Created Successfully")
}
