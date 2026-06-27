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
	initializers.DB.Migrator().DropTable(&users.User{})
	initializers.DB.Migrator().DropTable(&users.Role{})
	initializers.DB.Migrator().DropTable(&users.AccessRole{})

	// menus
	initializers.DB.Migrator().DropTable(&menus.Webmenu{})

	// masters
	initializers.DB.Migrator().DropTable(&masters.Goods{})

	// operations
	initializers.DB.Migrator().DropTable(&operations.Stock{})
	initializers.DB.Migrator().DropTable(&operations.Transaction{})

	fmt.Println("Database Deleted")
}
