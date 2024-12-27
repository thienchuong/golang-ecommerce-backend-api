package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/thienchuong/golang-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.PUT("/user/1", c.NewUserController().GetUserByID)
	}
	return r
}
