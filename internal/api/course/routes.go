package course

import "github.com/gin-gonic/gin"

func SetupCourseRoutes(r *gin.RouterGroup) {
	r.GET("/get/all/:page", getAllCoursesHandler)
	r.GET("/get/:id", getCourseHandler)

	r.POST("/create", createCourseHandler)

	r.PATCH("/update", updateCourseHandler)
	r.PATCH("/module/update", updateModuleHandler)
	r.PATCH("/lesson/update", updateLessonHandler)
	r.PATCH("/resource/update", updateResourceHandler)
}
