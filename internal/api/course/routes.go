package course

import "github.com/gin-gonic/gin"

func SetupCourseRoutes(r *gin.RouterGroup) {
	r.POST("/create", createCourseHandler)
}
