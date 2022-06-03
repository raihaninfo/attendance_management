package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/handlers"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/raihaninfo/attendance_magagment/docs"
	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := gin.Default()
	class := r.Group("/api")
	{
		class.GET("/", h.Home)
		class.GET("/class", h.Class)
		class.POST("/class", h.AddClass)
		class.PUT("/class/:id", h.UpdateClass)
		class.DELETE("/class/:id", h.DeleteClass)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(Port)
}
