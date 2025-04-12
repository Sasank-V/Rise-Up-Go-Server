package user

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	LearnerRole      Role = "learner"
	MentorRole       Role = "mentor"
	OrganisationRole Role = "organisation"
)

func CreateUserCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"email", "name", "role"},
		"properties": bson.M{
			"_id": bson.M{
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
	err := database.CreateCollection(db, lib.UserCollectionName, jsonSchema, []string{"_id", "email"})
	if err != nil {
		log.Fatal("Error creating user Collection: ", err)
		return
	}
	log.Println("User Collection Exists/Created Successfully")
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

	err := database.CreateCollection(db, lib.LearnerCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error creating Learner Collection: ", err)
	}
	log.Println("Learner Collection Exists/Created Successfully")
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

func CreateMentorCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"registered_courses": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"mentorship_requests": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"mentoship_sessions": bson.M{
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
			"reviews": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
			"available": bson.M{
				"bsonType": "bool",
			},
		},
	}
	err := database.CreateCollection(db, lib.MentorCollectionName, jsonSchema, []string{"user_id"})
	if err != nil {
		log.Fatal("Error Creating the Mentor Collection: ", err)
		return
	}
	log.Println("Mentor Collection Exists/Created Successfully")
}

func CreateEducationCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "institute", "degree", "start_date"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"institute": bson.M{
				"bsonType": "string",
			},
			"degree": bson.M{
				"bsonType": "string",
			},
			"start_date": bson.M{
				"bsonType": "date",
			},
			"end_date": bson.M{
				"bsonType": "date",
			},
		},
	}
	err := database.CreateCollection(db, lib.EducationCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error creating Education Collection: ", err)
		return
	}
	log.Println("Education Collection Exists/Created Successfully")
}

func CreateReviewCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"from", "to", "rating", "body"},
		"properties": bson.M{
			"from": bson.M{
				"bsonType": "string",
			},
			"to": bson.M{
				"bsonType": "string",
			},
			"rating": bson.M{
				"bsonType": "int",
				"minimum":  1,
				"maximum":  5,
			},
			"body": bson.M{
				"bsonType": "string",
			},
		},
	}
	err := database.CreateCollection(db, lib.ReviewCollectionName, jsonSchema, []string{})
	if err != nil {
		log.Fatal("Error Creating Review Collection: ", err)
		return
	}
	log.Println("Review Collection Exists/Created Successfully")
}

func CreateExperienceCollection(db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"user_id", "company", "position", "start_date"},
		"properties": bson.M{
			"user_id": bson.M{
				"bsonType": "string",
			},
			"company": bson.M{
				"bsonType": "string",
			},
			"position": bson.M{
				"bsonType": "string",
			},
			"start_date": bson.M{
				"bsonType": "date",
			},
			"end_date": bson.M{
				"bsonType": "date",
			},
			"description": bson.M{
				"bsonType": "string",
			},
		},
	}

	err := database.CreateCollection(db, lib.ExperienceCollectionName, jsonSchema, []string{"user_id", "company", "position"})
	if err != nil {
		log.Fatal("Error creating Experience Collection: ", err)
		return
	}
	log.Println("Experience Collection Exists/Created Successfully")
}

func ConnectAllUserCollections() {
	ConnectUserCollection()
	ConnectLearnerCollection()
	ConnectMentorCollection()
	ConnectOrganisationCollection()
	ConnectProfileCollections()
}
