package course

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/core/course"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/core/user"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/gin-gonic/gin"
)

func getAllCoursesHandler(c *gin.Context) {
	pageID := c.Param("page")
	id, err := strconv.ParseInt(pageID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing the pageId to int: %v", err),
		})
		return
	}
	courses, totalCount, err := course.GetAllCourses(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error getting all the courses: %v", err),
		})
		return
	}
	c.JSON(http.StatusAccepted, types.AllCoursesResponse{
		Page:       id,
		PageSize:   20,
		TotalCount: totalCount,
		Courses:    courses,
	})

}

func getCourseHandler(c *gin.Context) {
	courseID := c.Param("id")
	exists, err := course.CheckCourseExists(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error checking the course: %v", err),
		})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, types.MessageResponse{
			Message: "No course found with the given ID",
		})
		return
	}

	course, err := course.GetCoursewithID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error fetching course: %v", err),
		})
		return
	}

	c.JSON(http.StatusAccepted, course)
}

func createCourseHandler(c *gin.Context) {
	var info types.CreateCourseRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		log.Printf("Error Parsing Request Body: %v", err)
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}

	valid, err := user.CheckUserRole(info.UserID, string(user.OrganisationRole))
	if err != nil {
		log.Printf("Error checking user role: %v", err)
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: "Error checking user role",
		})
		return
	}

	if !valid {
		log.Printf("Not a Valid User to create Courses")
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: "Not a valid user to create courses",
		})
		return
	}

	err = course.AddCourse(info)
	if err != nil {
		log.Printf("Error creating course: %v", err)
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error creating course: %v", err),
		})
		return
	}

	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Course Created Successfully",
	})
}

func updateCourseHandler(c *gin.Context) {
	var info types.UpdateCourseRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing the request body: %v", err),
		})
		return
	}

	exists, err := course.CheckCourseExists(info.CourseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error checking course: %v", err),
		})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, types.MessageResponse{
			Message: fmt.Sprintf("No Course exists with the given ID"),
		})
		return
	}
	err = course.UpdateCourse(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error updating course: %v", err),
		})
		return
	}
	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Course Updated Successfully",
	})
}

func updateModuleHandler(c *gin.Context) {
	var info types.UpdateModule
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}

	exists, err := course.CheckModuleExists(info.ModuleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error checking module: %v", err),
		})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, types.MessageResponse{
			Message: "No module found with the given ID",
		})
		return
	}

	err = course.UpdateModule(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error updating module: %v", err),
		})
		return
	}

	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Module Updated Successfully",
	})
}

func updateLessonHandler(c *gin.Context) {
	var info types.UpdateLesson
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}
	exists, err := course.CheckLessonExists(info.LessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error checking lesson: %v", err),
		})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, types.MessageResponse{
			Message: "No course found with the given ID",
		})
		return
	}

	err = course.UpdateLesson(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error updating lesson: %v", err),
		})
		return
	}
	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Course Updated Successfully",
	})
}

func updateResourceHandler(c *gin.Context) {
	var info types.UpdateResource
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}

	exists, err := course.CheckResourceExists(info.ResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error checking resource: %v", err),
		})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, types.MessageResponse{
			Message: "No resource found with the given ID",
		})
		return
	}

	err = course.UpdateResource(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error updating resource: %v", err),
		})
		return
	}

	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Resource Updated Successfully",
	})
}
