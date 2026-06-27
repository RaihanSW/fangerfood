package masters

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Title string
	Body  string
}
