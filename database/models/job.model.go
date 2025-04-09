package models

import (
	"log"
	"time"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type WorkMode string

const (
	Remote WorkMode = "remote"
	OnSite WorkMode = "onsite"
	Hybrid WorkMode = "hybrid"
)

type JobType string

const (
	FullTime   JobType = "fulltime"
	PartTime   JobType = "parttime"
	Internship JobType = "intership"
)

type Job struct {
	Owner            string    `bson:"owner" json:"owner"`
	Title            string    `bson:"title" json:"title"`
	Description      string    `bson:"description" json:"description"`
	SkillTags        []string  `bson:"skill_tags" json:"skill_tags"`
	WorkMode         WorkMode  `bson:"work_mode" json:"work_mode"`
	JobType          JobType   `bson:"job_type" json:"job_type"`
	Location         string    `bson:"location" json:"location"`
	SalaryRangeStart int64     `bson:"salary_range_start" json:"salary_range_start"`
	SalaryRangeEnd   int64     `bson:"salary_range_end" json:"salary_range_end"`
	PostedAt         time.Time `bson:"posted_at" json:"posted_at"`
}

func CreateJobCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{},
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
	log.Printf("Job Collection Created Successfully")
}
