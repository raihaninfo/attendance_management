package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) FrontHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", nil)
}

func (h handler) FrontTeacher(c *gin.Context) {
	c.HTML(http.StatusOK, "teachers.gohtml", nil)
}
