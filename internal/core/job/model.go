package job

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Remote WorkMode = "remote"
	OnSite WorkMode = "onsite"
	Hybrid WorkMode = "hybrid"
)

const (
	FullTime   JobType = "fulltime"
	PartTime   JobType = "parttime"
	Internship JobType = "internship"
)

func CreateJobCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"owner", "title", "description", "work_mode", "job_type", "posted_at"},
		"properties": bson.M{
			"owner": bson.M{
				"bsonType": "string",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"description": bson.M{
				"bsonType": "string",
			},
			"skill_tags": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"work_mode": bson.M{
				"bsonType": "string",
				"enum":     []string{string(Remote), string(OnSite), string(Hybrid)},
			},
			"job_type": bson.M{
				"bsonType": "string",
				"enum":     []string{string(FullTime), string(PartTime), string(Internship)},
			},
			"location": bson.M{
				"bsonType": "string",
			},
			"salary_range_start": bson.M{
				"bsonType": "long",
			},
			"salary_range_end": bson.M{
				"bsonType": "long",
			},
			"evaluation_criteria": bson.M{
				"bsonType": "string",
			},
			"active": bson.M{
				"bsonType": "bool",
			},
			"contact": bson.M{
				"bsonType": "string",
			},
			"posted_at": bson.M{
				"bsonType": "date",
			},
		},
	}
	err := database.CreateCollection(db, lib.JobCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Job Collection: ", err)
		return
	}
	log.Println("Job Collection Exists/Created Successfully")
}

const (
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
	Pending  ApplicationStatus = "pending"
)

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
			"test_result": bson.M{
				"bsonType": "string",
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
	log.Println("Job Application Collection Exists/Created Successfully")

}

func ConnectAllJobCollections() {
	ConnectJobApplicationCollection()
	ConnectJobCollection()
}
