package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thienchuong/golang-ecommerce-backend-api/internal/service"
	"github.com/thienchuong/golang-ecommerce-backend-api/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc = user controller
// us = user service

// controller -> service -> repo -> models -> db
func (uc *UserController) GetUserByID(c *gin.Context) {
	err := 1
	if err != 1 {
		response.ErrorResponse(c, 20003, "no need")
		return
	}
	response.SuccessResponse(c, 20001, []string{"thienchuong-1", "thienchuong-2"})
}
