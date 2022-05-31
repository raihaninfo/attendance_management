package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raihaninfo/attendance_magagment/models"
	"github.com/raihaninfo/attendance_magagment/views"
)

var HomeView *views.View = views.NewView("views/frontend/home.gohtml")

func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	err := HomeView.Template.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
	// jj := models.Class{
	// 	ClassName:    "class one",
	// 	ClassTeacher: 2,
	// }
	// h.DB.Create(&jj)
	jj:= models.Class{}
	h.DB.First(&jj, 1)
	fmt.Println(jj)

}
