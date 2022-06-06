package models

type User struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	DateOfBirth string `json:"dob"`
	Designation string `json:"designation"`
	Password    string `json:"password"`
	UserType    int    `json:"type"`
	Status      int    `json:"status"`
}

type Student struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name    string `json:"name"`
	Roll    string `json:"roll"`
	Class   string `json:"class"`
	Section string `json:"section"`
	Gender  string `json:"gender"`
	Status  int    `json:"status"`
}

type Class struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	ClassName    string `json:"class"`
	ClassTeacher int    `json:"class-teacher" ` //gorm:"foreignKey:Id"
}

type Attendance struct {
	Id                int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name              string `json:"name"` //`json:"name" gorm:"foreignKey:Name"`
	StudentAttendance int    `json:"attendance"`
}
