package models

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	RoleId int
	Role   UserRole `json: "role" gorm`
	Points int      `json:"points"`
}

type UserRole struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
