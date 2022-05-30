package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name  string
	Email string
}

func Init() *gorm.DB {
	dbURL := "postgres://dev:secret@localhost/dev"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Test{})

	return db
}
