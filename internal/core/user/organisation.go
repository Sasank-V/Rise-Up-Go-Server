package user

import (
	"fmt"
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrganisationColl *mongo.Collection
var organisationConnect sync.Once

func ConnectOrganisationCollection() {
	organisationConnect.Do(func() {
		db := database.InitDB()
		CreateOrganisationCollection(db)
		OrganisationColl = db.Collection(lib.OrganisationCollectionName)
	})
}

func AddOrganisation(userID string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	newOrganisation := Organisation{
		UserID:             userID,
		Website:            "",
		JobsPosted:         []string{},
		CoursesPosted:      []string{},
		JobApplications:    []string{},
		MentorshipRequests: []string{},
	}

	res, err := OrganisationColl.InsertOne(ctx, newOrganisation)
	if err != nil {
		return "", err
	}

	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}
	return id, nil

}

func AddJobToOrganisation(orgUserID string, jobID string) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	user, err := GetBasicUserInfo(orgUserID)
	if err != nil {
		return err
	}
	orgID := user.RoleID

	orgobjID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": orgobjID,
	}
	update := bson.M{
		"$push": bson.M{
			"jobs_posted": jobID,
		},
	}
	res, err := OrganisationColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("no organisation found or jobId already exists")
	}
	return nil
}

func AddCourseToOrganisation(orgUserID string, courseID string) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	user, err := GetBasicUserInfo(orgUserID)
	if err != nil {
		return err
	}
	orgID := user.RoleID
	orgobjID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": orgobjID,
	}
	update := bson.M{
		"$push": bson.M{
			"courses_posted": courseID,
		},
	}
	res, err := OrganisationColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("no organisation found or courseId already exists")
	}
	return nil
}
