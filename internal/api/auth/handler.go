package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/core/user"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/gin-gonic/gin"
)

func signinHandler(c *gin.Context) {
	var info types.SigninRequest
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		log.Printf("Error parsing body: %v", err)
		c.JSON(http.StatusBadRequest, types.MessageResponse{
			Message: fmt.Sprintf("Error parsing request body: %v", err),
		})
		return
	}

	exists, err := user.UserExists(info.GoogleID)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		c.JSON(http.StatusInternalServerError, types.MessageResponse{
			Message: "Error checking User",
		})
		return
	}
	if !exists {
		err = user.AddUser(info)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			c.JSON(http.StatusInternalServerError, types.MessageResponse{
				Message: "Error creating new User",
			})
			return
		}
	}
	c.JSON(http.StatusOK, types.MessageResponse{
		Message: "User Signin Successfull",
	})
}
