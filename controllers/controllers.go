package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/attendance_magagment/handlers"
)

func Controller(Port string) {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("views/frontend/assets"))))

	r.HandleFunc("/", handlers.Home)
	http.ListenAndServe(Port, r)
}
