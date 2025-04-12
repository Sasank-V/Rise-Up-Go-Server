package course

import (
	"errors"
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var LessonColl *mongo.Collection
var connectLesson sync.Once

func ConnectLessonCollection() {
	connectLesson.Do(func() {
		db := database.InitDB()
		CreateLessonCollection(db)
		LessonColl = db.Collection(lib.LessonCollectionName)
	})
}

func CheckLessonExists(lessonID string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(lessonID)
	if err != nil {
		return false, err
	}

	err = LessonColl.FindOne(ctx, bson.M{
		"_id": objID,
	}).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func AddLesson(info types.CreateLesson, moduleID string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	newLesson := Lesson{
		ModuleID:    moduleID,
		Title:       info.Title,
		Description: info.Description,
		ContentLink: info.ContentLink,
		ContentType: ContentType(info.ContentType),
		Resources:   []string{},
		Duration:    info.Duration,
		OrderNo:     info.OrderNo,
	}

	res, err := LessonColl.InsertOne(ctx, newLesson)
	if err != nil {
		return "", err
	}
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}
	var resourceChan = make(chan string, len(info.Resources))
	var errChan = make(chan error, len(info.Resources))
	var wg sync.WaitGroup
	for _, resource := range info.Resources {
		wg.Add(1)
		go func(resource types.CreateResource, lessonID string) {
			defer wg.Done()
			resID, err := AddResource(resource, lessonID)
			if err != nil {
				// log.Println("Error Creating Resource: ", err)
				errChan <- err
				return
			}
			resourceChan <- resID
		}(resource, id)
	}

	go func() {
		wg.Wait()
		close(resourceChan)
		close(errChan)
		// log.Println("All Resources Created")
	}()

	for err := range errChan {
		return "", err
	}

	var resourceIDs []string
	for r := range resourceChan {
		resourceIDs = append(resourceIDs, r)
	}

	update := bson.M{
		"$set": bson.M{
			"resources": resourceIDs,
		},
	}

	_, err = LessonColl.UpdateByID(ctx, res.InsertedID, update)
	if err != nil {
		// log.Println("Error created Course: ", err)
		return "", err
	}

	// log.Println("Lesson Created: ", info.Title)
	return id, nil
}

func UpdateLesson(info types.UpdateLesson) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	updatedLesson := bson.M{}

	if info.Title != nil {
		updatedLesson["title"] = *info.Title
	}
	if info.Description != nil {
		updatedLesson["description"] = *info.Description
	}
	if info.ContentLink != nil {
		updatedLesson["content_link"] = *info.ContentLink
	}
	if info.ContentType != nil {
		updatedLesson["content_type"] = *info.ContentType
	}
	if info.Duration != nil {
		updatedLesson["duration"] = *info.Duration
	}
	if info.OrderNo != nil {
		updatedLesson["order_no"] = *info.OrderNo
	}

	if len(updatedLesson) == 0 {
		return errors.New("no fields to update")
	}

	objID, err := primitive.ObjectIDFromHex(info.LessonID)
	if err != nil {
		return errors.New("invalid lesson ID format")
	}

	res, err := LessonColl.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updatedLesson},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no lesson found to update or no fields were modified")
	}

	return nil
}
