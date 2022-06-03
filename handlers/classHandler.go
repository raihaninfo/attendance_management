package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/models"
)

func (h handler) AllClass(c *gin.Context) {
	var re []models.Class
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}

func (h handler) AddClass(c *gin.Context) {
	var class models.Class
	err := c.BindJSON(&class)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.DB.Create(&class)
	c.JSON(http.StatusOK, &class)
}

func (h handler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	if err := h.DB.Where("id = ?", id).First(&class).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&class)
	h.DB.Save(&class)
	c.JSON(http.StatusOK, &class)
}

func (h handler) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	err := h.DB.Where("id = ?", id).First(&class).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&class)
}
