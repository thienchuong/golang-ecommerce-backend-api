package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	udi := c.Query("udi")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"udi":     udi,
	})
}
