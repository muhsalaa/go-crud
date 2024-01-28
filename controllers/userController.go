package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
	"github.com/go-crud/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		utils.GenerateErrorResponse("Failed to Read Body", c)
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		utils.GenerateErrorResponse("Failed to Hash password", c)
		return
	}

	// create the user
	//  this string(hash) convert byte to string. String is composite component
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		utils.GenerateErrorResponse("Failed to Create User", c)
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		utils.GenerateErrorResponse("Failed to Read Body", c)
		return
	}

	// retrive user email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		utils.GenerateErrorResponse("Invalid Email or Password", c)
		return
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.GenerateErrorResponse("Invalid Email or Password", c)
		return
	}

	// generate jwt
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		utils.GenerateErrorResponse("Failed to create token", c)
		return
	}

	// set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	// receive user data from middleware
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}
