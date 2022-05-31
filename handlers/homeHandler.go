package handlers

import (
	"log"
	"net/http"

	"github.com/raihaninfo/attendance_magagment/views"
)

var HomeView *views.View = views.NewView("views/frontend/home.gohtml")

func Home(w http.ResponseWriter, r *http.Request) {
	err := HomeView.Template.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
