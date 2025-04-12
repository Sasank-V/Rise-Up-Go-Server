package user

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

var EducationColl *mongo.Collection
var ExperienceColl *mongo.Collection
var ReviewColl *mongo.Collection

var profileConnect sync.Once

func ConnectProfileCollections() {
	profileConnect.Do(func() {
		db := database.InitDB()
		CreateEducationCollection(db)
		CreateExperienceCollection(db)
		CreateReviewCollection(db)
		EducationColl = db.Collection(lib.EducationCollectionName)
		ExperienceColl = db.Collection(lib.ExperienceCollectionName)
		ReviewColl = db.Collection(lib.ReviewCollectionName)
	})
}
