package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
