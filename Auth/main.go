package main

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = make(map[string]*User)

// Global variable to store the JWT secret
var jwtSecret = []byte("Your_jwt_secret")

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
}

func handleRegister(c *gin.Context) {
	var newUser User

	err := c.BindJSON(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
		return
	}

	// Successfull registration
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Server Error"})
	}

	newUser.Password = string(hashedPassword)
	users[newUser.Email] = &newUser

	c.JSON(http.StatusOK, gin.H{"message": "Registration Successfull"})
}

func handleLogin(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Login failed"})
		return
	}

	// Check if user exists
	existingUser, ok := users[user.Email]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// check if entered password mathces
	passwordComp := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if passwordComp != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incorrect password"})
		return
	}

	// Successful Login

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email":   existingUser.Email,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": jwtToken})
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// check if Auth_Header is not empty
		authHeader := c.GetHeader("Authorization") // EXPECTED:- 'bearer':[token]
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		// check if Auth_Header is valid
		authParts := strings.Split(authHeader, " ")
		if len(authHeader) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.GET("/", welcomeHandler)
	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)

	router.GET("/secure", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a secure route"})
	})

	router.Run()
}
