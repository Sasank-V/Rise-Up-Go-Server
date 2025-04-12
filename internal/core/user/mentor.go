package user

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var MentorColl *mongo.Collection
var mentorConnect sync.Once

func ConnectMentorCollection() {
	mentorConnect.Do(func() {
		db := database.InitDB()
		CreateMentorCollection(db)
		MentorColl = db.Collection(lib.MentorCollectionName)
	})
}

func AddMentor(userID string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	newMentor := Mentor{
		UserID:             userID,
		Skills:             []string{},
		Experience:         []string{},
		RegisteredCourses:  []string{},
		MentorshipRequests: []string{},
		MentorShipSessions: []string{},
		TestsTaken:         []string{},
		Reviews:            []string{},
		Available:          true,
	}

	res, err := MentorColl.InsertOne(ctx, newMentor)
	if err != nil {
		return "", err
	}
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}
	return id, nil
}
