package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sasank-V/Rise-Up-Go-Server/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/database/models"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading Environment Variables")
	}
	fmt.Printf("Environment Variables Loaded")

	db := database.InitDB()

	models.CreateUserCollection(db)
	models.CreateLearnerCollection(db)
	models.CreateMentorCollection(db)
	

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello,World",
		})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
