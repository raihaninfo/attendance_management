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

	// json.Unmarshal()
	var abc []Abc
	json.Unmarshal(body, &abc)
	for _, i:= range abc{
		fmt.Println(i.Email)
	}
	c.HTML(http.StatusOK, "teachers.gohtml", abc)
}
