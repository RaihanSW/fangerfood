package operations

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Title string
	Body  string
}

type Transaction struct {
	gorm.Model
	Title string
	Body  string
}
