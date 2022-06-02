package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/handlers"
	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	router := gin.Default()
	// r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("views/frontend/assets"))))

	// r.HandleFunc("/", h.Home)
	router.GET("/ping", h.Home)
	router.GET("/abc", h.Teachers)

	// r.HandleFunc("/teacher", h.Teachers)
	// http.ListenAndServe(Port, r)
	router.Run(Port)
}
