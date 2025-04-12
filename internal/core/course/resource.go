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

var ResourceColl *mongo.Collection
var connectResource sync.Once

func ConnectResourceCollection() {
	connectResource.Do(func() {
		db := database.InitDB()
		CreateResourceCollection(db)
		ResourceColl = db.Collection(lib.ResourceCollectionName)
	})
}

func CheckResourceExists(resourceID string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(resourceID)
	if err != nil {
		return false, err
	}

	err = ResourceColl.FindOne(ctx, bson.M{"_id": objID}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func AddResource(info types.CreateResource, lessonID string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	newResource := Resource{
		LessonID: lessonID,
		Name:     info.Name,
		Link:     info.Link,
	}

	res, err := ResourceColl.InsertOne(ctx, newResource)
	if err != nil {
		return "", err
	}
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}
	// log.Println("Resource Created: ", info.Name)
	return id, nil
}

func UpdateResource(info types.UpdateResource) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	updatedResource := bson.M{}

	if info.Name != nil {
		updatedResource["name"] = *info.Name
	}
	if info.Link != nil {
		updatedResource["link"] = *info.Link
	}

	if len(updatedResource) == 0 {
		return errors.New("no fields to update")
	}

	objID, err := primitive.ObjectIDFromHex(info.ResourceID)
	if err != nil {
		return errors.New("invalid resource ID format")
	}

	res, err := ResourceColl.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updatedResource},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no resource found to update or no fields were modified")
	}

	return nil
}
