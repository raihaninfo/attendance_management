package models

type FrontStudent struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name    string `json:"name"`
	Roll    string `json:"roll"`
	Class   string `json:"class"`
	Section string `json:"section"`
	Gender  string `json:"gender"`
	Status  int    `json:"status"`
}
