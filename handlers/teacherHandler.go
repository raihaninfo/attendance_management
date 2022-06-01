package handlers

import (
	"log"
	"net/http"

	"github.com/raihaninfo/attendance_magagment/views"
)

var TeacherView *views.View = views.NewView("views/frontend/teachers.gohtml")

func (h handler) Teachers(w http.ResponseWriter, r *http.Request) {
	err := TeacherView.Template.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
