package user_hdl

import (
	"loyalty/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	userService ports.UserService
}

func New(userService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := hdl.userService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}
