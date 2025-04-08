package models

import (
	"log"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	GoogleID string `bson:"google_id" json:"google_id"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Picture  string `bson:"picture" json:"picture"`
	Role     string `bson:"role" json:"role"`
	RoleID   string `bson:"role_id" json:"role_id"`
}

func CreateUserCollection(db *mongo.Database) {
	ctx, cancel := database.GetContext()
	defer cancel()

	exists, err := database.CollectionExist(db, lib.UserCollectionName)
	if err != nil {
		log.Fatal("Error checkiing the existing Collection: ", err)
		return
	}
	if exists {
		log.Printf("User Collection Already Exists, Skipping Creation...\n")
		return
	}
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
			"role": bson.M{
				"bsonType": "string",
			},
			"role_id": bson.M{
				"bsonType": "string",
			},
		},
	}
	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := options.CreateCollection().SetValidator(validator)
	err = db.CreateCollection(ctx, lib.UserCollectionName, opts)
	if err != nil {
		log.Fatal("Error creating User Collection ", err)
		return
	}
	if err = database.SetUniqueKeys(db.Collection(lib.UserCollectionName), []string{"google_id", "email"}); err != nil {
		log.Fatal("Error setting up Unique Keys ", err)
		return
	}
}
