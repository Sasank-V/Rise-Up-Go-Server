package course

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/gin-gonic/gin"
)

func createCourseHandler(c *gin.Context) {
	var info types.CreateCourseRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		log.Printf("Error Parsing Request Body: %v", err)
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}
	
}
