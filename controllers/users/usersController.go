package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"fangerfood/initializers"
	"fangerfood/models/users"
)

func Signup(ctx *gin.Context) {
	// Ambil data jsonnya
	var body struct {
		Username string
		Password string
		RoleId   int
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch body data",
		})

		return
	}

	// Hash passwordnya
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password, check the password given",
		})

		return
	}
	// Create usernya, input ke db
	knownUser := users.User{
		Username: body.Username,
		Password: string(hash),
	}
	result := initializers.DB.Create(&knownUser)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User",
		})

		return
	}

	// kasih response json
	ctx.JSON(http.StatusOK, gin.H{"data": knownUser})

}

func Login(ctx *gin.Context) {
	// Ambil data jsonnya
	var body struct {
		Username string
		Password string
		RoleId   int
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch json body data",
		})

		return
	}

	// Look up requested User
	var user users.User
	initializers.DB.First(&user, "Username = ?", body.Username)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No user with this Username",
		})

		return
	}
	// compare password hashing
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Password",
		})

		return
	}
	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretT0oken"))
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong creating token",
		})

		return
	}

	// kasih response json
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"token": tokenString,
	}})

}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("User")

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"UserData": user,
	}})
}
