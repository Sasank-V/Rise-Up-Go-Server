package job

import (
	"fmt"
	"net/http"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/core/job"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/gin-gonic/gin"
)

func createJobHandler(c *gin.Context) {
	var info types.CreateJobRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing the request body: %v", err),
		})
		return
	}
	err := job.AddJob(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error creating job: %v", err),
		})
		return
	}
	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Job Created Successfully",
	})
}

func updateJobHandler(c *gin.Context) {
	var info types.UpdateJobRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}

	err := job.UpdateJob(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: fmt.Sprintf("Error updating job: %v", err),
		})
		return
	}
	c.JSON(http.StatusAccepted, types.MessageResponse{
		Message: "Job Update Successfully",
	})
}
