package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	r.POST("/signin", signinHandler)
}

func signinHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Signin Route",
	})
}
