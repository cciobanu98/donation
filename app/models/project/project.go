package project

// Project struct
type Project struct {
	ID    uint   	`json:"id" gorm:"primary_key"`
	Title string 	`json:"title"`
}

