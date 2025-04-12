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

var ModuleColl *mongo.Collection
var connectModule sync.Once

func ConnectModuleCollection() {
	connectModule.Do(func() {
		db := database.InitDB()
		CreateModuleCollection(db)
		ModuleColl = db.Collection(lib.ModuleCollectionName)
	})
}

func CheckModuleExists(moduleID string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(moduleID)
	if err != nil {
		return false, err
	}

	err = ModuleColl.FindOne(ctx, bson.M{"_id": objID}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func AddModule(info types.CreateModule, courseId string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()
	newModule := Module{
		CourseID: courseId,
		Title:    info.Title,
		OrderNo:  info.OrderNo,
		Lessons:  []string{},
	}
	res, err := ModuleColl.InsertOne(ctx, newModule)
	if err != nil {
		return "", err
	}
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}

	var lessonChan = make(chan string, len(info.Lessons))
	var errChan = make(chan error, len(info.Lessons))
	var wg sync.WaitGroup

	for _, lesson := range info.Lessons {
		wg.Add(1)
		go func(lesson types.CreateLesson, moduleId string) {
			defer wg.Done()
			lessonID, err := AddLesson(lesson, moduleId)
			if err != nil {
				// log.Println("Error Creating Lesson: ", err)
				errChan <- err
				return
			}
			lessonChan <- lessonID
		}(lesson, id)
	}

	go func() {
		wg.Wait()
		close(lessonChan)
		close(errChan)
		// log.Println("All Lessons Created")

	}()

	for err := range errChan {
		return "", err
	}

	var lessonIDs []string
	for l := range lessonChan {
		lessonIDs = append(lessonIDs, l)
	}

	update := bson.M{
		"$set": bson.M{
			"lessons": lessonIDs,
		},
	}
	_, err = LessonColl.UpdateByID(ctx, res.InsertedID, update)
	if err != nil {
		return "", err
	}
	// log.Println("Module Created: ", info.Title)
	return id, nil
}

func UpdateModule(info types.UpdateModule) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	updatedModule := bson.M{}

	if info.Title != nil {
		updatedModule["title"] = *info.Title
	}
	if info.OrderNo != nil {
		updatedModule["order_no"] = *info.OrderNo
	}

	if len(updatedModule) == 0 {
		return errors.New("no fields to update")
	}

	objID, err := primitive.ObjectIDFromHex(info.ModuleID)
	if err != nil {
		return errors.New("invalid module ID format")
	}

	res, err := ModuleColl.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updatedModule},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no module found to update or no fields were modified")
	}

	return nil
}
