package user

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var LearnerColl *mongo.Collection
var learnerConnect sync.Once

func ConnectLearnerCollection() {
	learnerConnect.Do(func() {
		db := database.InitDB()
		CreateLearnerCollection(db)
		LearnerColl = db.Collection(lib.LearnerCollectionName)
	})
}

func AddLearner(userID string) (string, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	newLearner := Learner{
		UserID:             userID,
		Skills:             []string{},
		Interests:          []string{},
		JobPreferences:     []string{},
		LanguagePreferred:  "en",
		Education:          []string{},
		ProfileCompletion:  10,
		EnrolledCourses:    []string{},
		AppliedJobs:        []string{},
		TestsTaken:         []string{},
		MentorshipRequests: []string{},
		MentorshipSessions: []string{},
		Reviews:            []string{},
	}

	res, err := LearnerColl.InsertOne(ctx, newLearner)
	if err != nil {
		return "", err
	}

	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return "", err
	}
	return id, nil
}
