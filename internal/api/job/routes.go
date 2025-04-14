package job

import "github.com/gin-gonic/gin"

func SetUpJobRoutes(r *gin.RouterGroup) {
	r.POST("/create", createJobHandler)
	r.PATCH("/update", updateJobHandler)
}
