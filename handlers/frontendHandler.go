package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) FrontHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", nil)
}

type Abc struct {
	Id          int
	Name        string
	UserName    string
	Email       string
	Mobile      string
	DateOfBirth string
	Designation string
	Password    string
	UserType    int
	Status      int
}

var abc Abc

func (h handler) FrontTeacher(c *gin.Context) {
	req, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &abc)
	fmt.Println(abc)
	c.HTML(http.StatusOK, "teachers.gohtml", gin.H{
		"Name":  abc.Name,
		"Email": "raihanmahmudi35@gmail.com",
	})
}
