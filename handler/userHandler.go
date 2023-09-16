package handler

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(loginRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}
