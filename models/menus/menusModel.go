package menus

import "gorm.io/gorm"

type Webmenu struct {
	gorm.Model
	Title string
	Body  string
}
