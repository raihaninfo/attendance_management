package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/attendance_magagment/handlers"
	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("views/frontend/assets"))))

	r.HandleFunc("/", h.Home)
	http.ListenAndServe(Port, r)
}
