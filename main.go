package main

import (
	"fangerfood/initializers"

	"github.com/gin-gonic/gin"

	//routes list
	UserRoutes "fangerfood/routes/users"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	UserRoutes.UserRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
