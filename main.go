package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	auth_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/auth"
	course_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/course"
	job_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/job"
	mentorship_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/mentorship"
	test_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/test"
	user_api "github.com/Sasank-V/Rise-Up-Go-Server/internal/api/user"
	"github.com/gin-contrib/cors"

	course_core "github.com/Sasank-V/Rise-Up-Go-Server/internal/core/course"
	job_core "github.com/Sasank-V/Rise-Up-Go-Server/internal/core/job"
	mentorship_core "github.com/Sasank-V/Rise-Up-Go-Server/internal/core/mentorship"
	test_core "github.com/Sasank-V/Rise-Up-Go-Server/internal/core/test"
	user_core "github.com/Sasank-V/Rise-Up-Go-Server/internal/core/user"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func CreateAllCollections() {
	db := database.InitDB()
	user_core.ConnectAllUserCollections()
	course_core.ConnectAllCourseCollections()
	job_core.ConnectAllJobCollections()
	test_core.CreateAllTestCollections(db)
	mentorship_core.CreateAllMentorshipCollections(db)
}

func SetupRoutes(r *gin.Engine) {
	authApi := r.Group("/api/auth")
	userApi := r.Group("/api/user")
	courseApi := r.Group("/api/course")
	jobApi := r.Group("/api/job")
	mentorshipApi := r.Group("/api/mentorship")
	testApi := r.Group("/api/test")

	auth_api.SetupAuthRoutes(authApi)
	user_api.SetupUserRoutes(userApi)
	course_api.SetupCourseRoutes(courseApi)
	job_api.SetUpJobRoutes(jobApi)
	mentorship_api.SetupMentorshipRoutes(mentorshipApi)
	test_api.SetupTestRoutes(testApi)
}

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading Environment Variables")
		return
	}
	fmt.Println("Environment Variables Loaded")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	CreateAllCollections()
	SetupRoutes(r)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello,World",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
