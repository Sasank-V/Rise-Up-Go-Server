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
	Skills             []string `bson:"skills" json:"skills"`
	Interests          []string `bson:"interests" json:"interests"`
	JobPreferences     []string `bson:"job_preferences" json:"job_preferences"` //Later
	LanguagePreferred  string   `bson:"language_preferred" json:"language_preferred"`
	Education          []string `bson:"education" json:"education"`
	ProfileCompletion  int      `bson:"profile_completion" json:"profile_completion"`
	EnrolledCourses    []string `bson:"enrolled_courses" json:"enrolled_courses"`
	AppliedJobs        []string `bson:"applied_jobs" json:"applied_jobs"`
	TestsTaken         []string `bson:"tests_taken" json:"tests_taken"`
	MentorshipRequests []string `bson:"mentorship_requests" json:"mentorship_requests"`
	MentorshipSessions []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
	Reviews            []string `bson:"reviews" json:"reviews"`
}

func CreateLearnerCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"skills": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"interests": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"job_preferences": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"language_preferred": bson.M{
				"bsonType": "string",
			},
			"education": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"enrolled_courses": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"applied_jobs": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"tests_taken": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"mentorship_requests": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"mentorship_sessions": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"reviews": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"profile_completion": bson.M{
				"bsonType": "int",
			},
		},
	}

	err := database.CreateCollection(db, lib.LeanerCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error creating Learner Collection: ", err)
	}
	log.Printf("Learner Collection Exists/Created Successfully")
}
