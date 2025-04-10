package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Organisation struct {
	UserID             string   `bson:"user_id" json:"user_id"`
	OrganisationName   string   `bson:"organisation_name" json:"organisation_name"`
	About              string   `bson:"about" json:"about"`
	Website            string   `bson:"website" json:"website"`
	JobsPosted         []string `bson:"jobs_posted" json:"jobs_posted"`
	CoursesPosted      []string `bson:"courses_posted" json:"course_posted"`
	JobApplications    []string `bson:"job_applications" json:"job_applications"`
	MentorshipRequests []string `bson:"mentorship_sessions" json:"mentorship_sessions"`
}

func CreateOrganisationCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"organisation_name": bson.M{
				"bsonType": "string",
			},
			"about": bson.M{
				"bsonType": "string",
			},
			"website": bson.M{
				"bsonType": "string",
			},
			"jobs_posted": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"courses_posted": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"job_applications": bson.M{
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

	err := database.CreateCollection(db, lib.OrganisationCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error creating Organisation Collection: ", err)
		return
	}
	log.Println("Organisation Collection Exists/Created Successfully")
}
