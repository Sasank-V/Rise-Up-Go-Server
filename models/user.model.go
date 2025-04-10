package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Role string

const (
	LearnerRole      Role = "learner"
	MentorRole       Role = "mentor"
	OrganisationRole Role = "organisation"
)

type User struct {
	GoogleID     string `bson:"google_id" json:"google_id"`
	Name         string `bson:"name" json:"name"`
	Email        string `bson:"email" json:"email"`
	Picture      string `bson:"picture" json:"picture"`
	Bio          string `bson:"bio" json:"bio"`
	Location     string `bson:"location" json:"location"`
	Role         string `bson:"role" json:"role"`
	RoleID       string `bson:"role_id" json:"role_id"`
	AccessToken  string `bson:"access_token" json:"access_token"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
	ExpiresAt    string `bson:"expires_at" json:"expires_at"`
}

func CreateUserCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"google_id", "email", "name", "role"},
		"properties": bson.M{
			"google_id": bson.M{
				"bsonType": "string",
			},
			"name": bson.M{
				"bsonType": "string",
			},
			"email": bson.M{
				"bsonType": "string",
			},
			"picture": bson.M{
				"bsonType": "string",
			},
			"bio": bson.M{
				"bsonType": "string",
			},
			"location": bson.M{
				"bsonType": "string",
			},
			"role": bson.M{
				"bsonType": "string",
				"enum":     []string{string(LearnerRole), string(MentorRole), string(OrganisationRole)},
			},
			"role_id": bson.M{
				"bsonType": "string",
			},
			"access_token": bson.M{
				"bsonType": "string",
			},
			"refresh_token": bson.M{
				"bsonType": "string",
			},
			"expires_at": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.UserCollectionName, jsonSchema, []string{"google_id", "email"})
	if err != nil {
		log.Fatal("Error creating user Collection: ", err)
		return
	}
	log.Printf("User Collection Exists/Created Successfully\n")
}
