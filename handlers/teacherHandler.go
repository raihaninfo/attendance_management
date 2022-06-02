package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/models"
)

func (h handler) Teacher(c *gin.Context) {
	var re []models.Class
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}
