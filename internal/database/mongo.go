package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBClient   *mongo.Client
	clientOnce sync.Once
)

func InitDB() *mongo.Database {
	connectDB()
	DBName := os.Getenv("DATABASE_NAME")
	if DBName == "" {
		log.Fatal("Set your 'DATABASE_NAME' environment variable")
	}
	return DBClient.Database(DBName)
}

func connectDB() {
	clientOnce.Do(func() {
		uri := os.Getenv("CONNECTION_STRING")
		if uri == "" {
			log.Fatal("Set your 'CONNECTION_STRING' environment variable")
		}
		dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

		if err != nil {
			log.Fatal("[MONGO-DB] Error connecting database", err)
		}
		if err := dbClient.Ping(context.TODO(), nil); err != nil {
			log.Fatal("[MONGO-DB] Database connection test failed", err)
		}
		fmt.Println("[MONGO-DB] Database Connected Successfully")
		DBClient = dbClient
	})
}

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func SetUniqueKeys(coll *mongo.Collection, fields []string) error {
	ctx, cancel := GetContext()
	defer cancel()

	for _, field := range fields {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{field, 1}},
			Options: options.Index().SetUnique(true),
		}
		_, err := coll.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			return err
		}
	}
	return nil
}

func CollectionExist(db *mongo.Database, coll_name string) (bool, error) {
	ctx, cancel := GetContext()
	defer cancel()

	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	if slices.Contains(collections, coll_name) {
		return true, nil
	}
	return false, nil
}

func CreateCollection(db *mongo.Database, coll_name string, coll_schema bson.M, unique_fields []string) error {
	ctx, cancel := GetContext()
	defer cancel()

	exists, err := CollectionExist(db, coll_name)
	if err != nil {
		return err
	}
	if exists {
		log.Printf("%v Collection Already Exists , Skipping Creation...\n", coll_name)
		return nil
	}

	validator := bson.M{
		"$jsonSchema": coll_schema,
	}
	opts := options.CreateCollection().SetValidator(validator)
	err = db.CreateCollection(ctx, coll_name, opts)
	if err != nil {
		return err
	}
	if err := SetUniqueKeys(db.Collection(coll_name), unique_fields); err != nil {
		return err
	}
	return nil
}
