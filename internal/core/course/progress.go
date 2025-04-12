package course

import (
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

var CourseProgressColl *mongo.Collection
var connectCourseProgress sync.Once

func ConnectCourseProgressCollection() {
	connectCourseProgress.Do(func() {
		db := database.InitDB()
		CreateCourseProgressCollection(db)
		CourseProgressColl = db.Collection(lib.CourseProgressCollectionName)
	})

}
