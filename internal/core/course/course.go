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

var CourseColl *mongo.Collection
var courseConnect sync.Once

func ConnectCourseCollection() {
	courseConnect.Do(func() {
		db := database.InitDB()
		CreateCourseCollection(db)
		CourseColl = db.Collection(lib.CourseCollectionName)
	})
}

func CheckCourseExists(courseID string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return false, err
	}

	err = CourseColl.FindOne(ctx, bson.M{"_id": objID}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func AddCourse(info types.CreateCourseRequest) error {
	ctx, cancel := database.GetContext()
	defer cancel()
	newCourse := Course{
		Owner:         info.UserID,
		Banner:        info.Banner,
		Title:         info.Title,
		Description:   info.Description,
		Difficulty:    Difficulty(info.Difficulty),
		Duration:      info.Duration,
		Skills:        info.Skills,
		Modules:       []string{},
		Instructors:   info.Instructors,
		Discussions:   []string{},
		Prerequisites: info.Prerequisites,
		Outcomes:      info.Outcomes,
	}

	res, err := CourseColl.InsertOne(ctx, newCourse)
	if err != nil {
		return err
	}
	// fmt.Printf("CourseID: %v", res.InsertedID)
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var moduleChan = make(chan string, len(info.Modules))
	var errChan = make(chan error, len(info.Modules))

	for _, module := range info.Modules {
		wg.Add(1)
		go func(module types.CreateModule, courseID string) {
			defer wg.Done()
			moduleID, err := AddModule(module, courseID)
			if err != nil {
				// log.Println("Error Creating Module: ", err)
				errChan <- err
				return
			}
			moduleChan <- moduleID
		}(module, id)
	}

	go func() {
		wg.Wait()
		close(moduleChan)
		close(errChan)
		// log.Println("All Modules Created")
	}()

	for err := range errChan {
		return err
	}

	var moduleIDs []string
	for m := range moduleChan {
		moduleIDs = append(moduleIDs, m)
	}

	update := bson.M{
		"$set": bson.M{
			"modules": moduleIDs,
		},
	}
	_, err = CourseColl.UpdateByID(ctx, res.InsertedID, update)
	if err != nil {
		return err
	}
	// log.Println("Course Created: ", info.Title)
	return nil
}

func UpdateCourse(info types.UpdateCourseRequest) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	updatedCourse := bson.M{}

	if info.Title != nil {
		updatedCourse["title"] = *info.Title
	}
	if info.Banner != nil {
		updatedCourse["banner"] = *info.Banner
	}
	if info.Description != nil {
		updatedCourse["description"] = *info.Description
	}
	if info.Difficulty != nil {
		updatedCourse["difficulty"] = *info.Difficulty
	}
	if info.Duration != nil {
		updatedCourse["duration"] = *info.Duration
	}
	if info.Skills != nil {
		updatedCourse["skills"] = info.Skills
	}
	if info.Prerequisites != nil {
		updatedCourse["prerequisites"] = *info.Prerequisites
	}
	if info.Outcomes != nil {
		updatedCourse["outcomes"] = *info.Outcomes
	}

	if len(updatedCourse) == 0 {
		return errors.New("no fields to update")
	}

	objID, err := primitive.ObjectIDFromHex(info.CourseID)
	if err != nil {
		return errors.New("invalid course ID format")
	}

	res, err := CourseColl.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updatedCourse},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no course found to update or no fields were modified")
	}

	return nil
}
