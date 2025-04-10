package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/models"
	"github.com/Sasank-V/Rise-Up-Go-Server/models/course"
	"github.com/Sasank-V/Rise-Up-Go-Server/models/job"
	"github.com/Sasank-V/Rise-Up-Go-Server/models/mentorship"
	"github.com/Sasank-V/Rise-Up-Go-Server/models/profile"
	"github.com/Sasank-V/Rise-Up-Go-Server/models/test"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func CreateAllCollections() {
	db := database.InitDB()

	models.CreateUserCollection(db)
	models.CreateLearnerCollection(db)
	models.CreateMentorCollection(db)
	models.CreateOrganisationCollection(db)

	profile.CreateEducationCollection(db)
	profile.CreateExperienceCollection(db)
	profile.CreateReviewCollection(db)

	course.CreateCourseCollection(db)
	course.CreateCourseProgressCollection(db)
	course.CreateLessonCollection(db)
	course.CreateModuleCollection(db)
	course.CreateResourceCollection(db)

	job.CreateJobCollection(db)
	job.CreateJobApplicationCollection(db)

	mentorship.CreateMentorShipRequestCollection(db)
	mentorship.CreateMentorShipSessionCollection(db)

	test.CreateTestCollection(db)
	test.CreateTestResultCollection(db)
}

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading Environment Variables")
	}
	fmt.Printf("Environment Variables Loaded")

	CreateAllCollections()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello,World",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
