package db

import (
	"log"

	"github.com/raihaninfo/attendance_magagment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://dev:secret@localhost/dev"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{}, &models.Class{}, &models.Student{})

	return db
}
