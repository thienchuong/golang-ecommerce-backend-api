package routers

import (
	"net/http"

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

func pong(c *gin.Context) {
	name := c.Param("name")
	udi := c.Query("udi")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"udi":     udi,
	})
}
