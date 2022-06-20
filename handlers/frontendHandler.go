package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/attendance_magagment/models"
)

func (h handler) FrontHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", nil)
}

func (h handler) FrontTeacher(c *gin.Context) {
	req, err := http.Get("http://localhost:8082/api/user")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var teacher []models.Teacher
	json.Unmarshal(body, &teacher)

	c.HTML(http.StatusOK, "teachers.gohtml", teacher)
}

func (h handler) FrontStudent(c *gin.Context) {
	req, err := http.Get("http://localhost:8082/api/student")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var student []models.FrontStudent
	json.Unmarshal(body, &student)
	
	c.HTML(http.StatusOK, "student.gohtml", student)
}
