package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/raihaninfo/attendance_magagment/docs"
	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", h.Home)
		api.GET("/teacher", h.Teachers)
		api.POST("/teacher")
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(Port)
}
