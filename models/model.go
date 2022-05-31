package models

type User struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	UserName string
	Email    string
	Mobile   string
	Password string
	UserType int
	Status   int
}

type Student struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Roll   string
	Class  string
	Status int
}

type Class struct {
	Id           int `gorm:"primaryKey"`
	ClassName    string
	ClassTeacher User
}
