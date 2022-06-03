package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/models"
)

func (h handler) AllStudent(c *gin.Context) {
	var re []models.Student
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}

func (h handler) AddStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.DB.Create(&student)
	c.JSON(http.StatusOK, &student)
}

func (h handler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := h.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&student)
	h.DB.Save(&student)
	c.JSON(http.StatusOK, &student)
}

func (h handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	err := h.DB.Where("id = ?", id).First(&student).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&student)
}
