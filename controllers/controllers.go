package controllers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/handlers"

	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := gin.Default()
	r.GET("/api", h.Home)

	r.Use(static.Serve("/asset", static.LocalFile("views/frontend/assets", true)))
	r.LoadHTMLGlob("views/frontend/*.gohtml")

	class := r.Group("/api")
	{
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

	user := r.Group("/api")
	{
		user.GET("/user", h.AllUser)
		user.POST("/user", h.AddUser)
		user.PUT("/user/:id", h.UpdateUser)
		user.DELETE("/user/:id", h.DeleteUser)
	}

	attendance := r.Group("/api")
	{
		attendance.GET("/attendance", h.Attendance)
	}

	// frontend
	frontend := r.Group("/")
	{
		frontend.GET("/", h.FrontHome)
		frontend.GET("/teacher", h.FrontTeacher)
		frontend.GET("/student", h.FrontStudent)
	}

	r.Run(Port)
}
