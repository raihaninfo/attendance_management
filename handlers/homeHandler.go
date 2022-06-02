package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Home(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"App":     "Attendance Management system",
		"Version": "1.0",
		"Author":  "Md Abu Raihan",
		"Url":     "https://github.com/raihaninfo",
	})

}
