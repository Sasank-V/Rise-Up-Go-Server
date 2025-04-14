package job

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

var JobApplicationColl *mongo.Collection
var connectJobApplication sync.Once

func ConnectJobApplicationCollection() {
	connectJobApplication.Do(func() {
		db := database.InitDB()
		CreateJobApplicationCollection(db)
		JobApplicationColl = db.Collection(lib.JobApplicationCollectionName)
	})
}
