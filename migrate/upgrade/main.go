package main

import (
	"fangerfood/initializers"
	"fangerfood/models/masters"
	"fangerfood/models/menus"
	"fangerfood/models/operations"
	"fangerfood/models/users"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// users
	initializers.DB.AutoMigrate(&users.User{})
	initializers.DB.AutoMigrate(&users.Role{})
	initializers.DB.AutoMigrate(&users.AccessRole{})

	// menus
	initializers.DB.AutoMigrate(&menus.Webmenu{})

	// masters
	initializers.DB.AutoMigrate(&masters.Goods{})

	// operations
	initializers.DB.AutoMigrate(&operations.Stock{})
	initializers.DB.AutoMigrate(&operations.Transaction{})

	fmt.Println("Database Created")
}
