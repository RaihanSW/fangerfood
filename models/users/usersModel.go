package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password  string
	RoleId int
}

type Role struct {
	gorm.Model
	Name string
}

type AccessRole struct {
	gorm.Model
	Title string
	Body  string
}
