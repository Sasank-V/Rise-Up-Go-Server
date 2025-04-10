package job

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ApplicationStatus string

const (
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
	Pending  ApplicationStatus = "pending"
)

type JobApplication struct {
	UserID          string            `bson:"user_id" json:"user_id"`
	JobID           string            `bson:"job_id" json:"job_id"`
	Status          ApplicationStatus `bson:"status" json:"status"`
	MatchPercentage int               `bson:"match_percentage" json:"match_percentage"`
}

func CreateJobApplicationCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "job_id", "status", "match_percentage"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"job_id": bson.M{
				"bsonType": "string",
			},
			"status": bson.M{
				"bsonType": "string",
				"enum":     []string{string(Accepted), string(Rejected), string(Pending)},
			},
			"match_percentage": bson.M{
				"bsonType": "int",
				"minimum":  0,
				"maximum":  100,
			},
		},
	}

	err := database.CreateCollection(db, lib.JobApplicationCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Job Application Collection: ", err)
		return
	}
	log.Printf("Job Application Collection Exists/Created Successfully")

}
