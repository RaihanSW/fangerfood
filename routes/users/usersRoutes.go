package users

import (
	UserController "fangerfood/controllers/users"
	"fangerfood/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	router.POST("/users/signup", UserController.Signup)
	router.POST("/users/login", UserController.Login)
	router.GET("/users/validate", middleware.RequireAuth, UserController.Validate)
}
