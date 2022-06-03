package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/handlers"

	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := gin.Default()
	class := r.Group("/api")
	{
		class.GET("/", h.Home)
		class.GET("/class", h.AllClass)
		class.POST("/class", h.AddClass)
		class.PUT("/class/:id", h.UpdateClass)
		class.DELETE("/class/:id", h.DeleteClass)
	}
	student := r.Group("/api")
	{
		student.GET("/student", h.AllStudent)
		student.POST("/student", h.AddStudent)
		student.PUT("/student/:id", h.UpdateStudent)
		student.DELETE("/student/:id", h.DeleteStudent)
	}

	r.Run(Port)
}
