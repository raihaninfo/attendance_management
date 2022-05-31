package main

import (
	"github.com/raihaninfo/attendance_magagment/controllers"
)

var Port string = ":8082"

func main() {
	controllers.Controller(Port)
}
