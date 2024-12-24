package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", pong)
		v1.PUT("/ping", pong)
		v1.PATCH("/ping", pong)
		v1.DELETE("/ping", pong)
		v1.HEAD("/ping", pong)
		v1.OPTIONS("/ping", pong)
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
