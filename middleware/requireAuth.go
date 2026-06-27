package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"fangerfood/initializers"
	"fangerfood/models/users"
)

func RequireAuth(ctx *gin.Context) {
	// Get cookie of the request
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// validate/decorate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// find user with the token
		var user users.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to request
		ctx.Set("User", user)
	}

	// Continue
	ctx.Next()

}
