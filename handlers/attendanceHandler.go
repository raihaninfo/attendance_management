package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/models"
)

func (h handler) Attendance(c *gin.Context) {
	var attendance []models.Attendance
	h.DB.Find(&attendance)
	c.IndentedJSON(http.StatusOK, attendance)
}
