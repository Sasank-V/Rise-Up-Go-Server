package user

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
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
		OrganisationName:   "",
		About:              "",
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
