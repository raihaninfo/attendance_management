package main

import (
	"fmt"

	"github.com/raihaninfo/attendance_magagment/db"
)

func main() {
	db.Init()
	fmt.Println("Hello Bangladesh!")
}
