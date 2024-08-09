package controller

import (
	"Implementing_Authentication_Authorization/data"
	"Implementing_Authentication_Authorization/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *data.UserCollection

func init() {
	UserCollection = data.NewUserCollection()
}

func HandelUserRegister(c *gin.Context) {
	var curr_user models.User

	if err := c.ShouldBindJSON(&curr_user); err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(curr_user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "internal server error"})
	}

	curr_user.Password = string(hashedPassword)
	UserCollection.Users[curr_user.Email] = &curr_user

	c.IndentedJSON(200, gin.H{"message": "user registered successfully"})
}

func HandelUserLogin(c *gin.Context) {
	var curr_user models.User

	if err := c.ShouldBindJSON(&curr_user); err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid request"})
	}

	// fetch user
	existingUser, ok := UserCollection.Users[curr_user.Email]

	// check if curr_user.Email exists
	if !ok {
		c.IndentedJSON(401, gin.H{"error": "Invalid Email"})
	}

	// check if user has inputed the correct password
	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(curr_user.Password)) != nil {
		c.IndentedJSON(401, gin.H{"error": "Incorrect password"})
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email":   existingUser.Email,
	})

	// sign the token with the secret key 'jwtSecret'
	jwtToken, err := token.SignedString(data.JwtSecret)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "user logged in successfully", "token": jwtToken})
}

func HandelSecureMessage(c *gin.Context) {
	// retrieve claims from the context
	claimsValue, exists := c.Get("claims")

	if !exists {
		c.IndentedJSON(500, gin.H{"error": "no claims found"})
		c.Abort()
		return
	}

	// retrieve jwt.MapClaims from the claimsValue
	claims, ok := claimsValue.(jwt.MapClaims)
	if !ok {
		c.IndentedJSON(500, gin.H{"error": "Claims are not valid"})
		c.Abort()
		return
	}

	// the two claims used to instantiate the JWT -> user_id and email

	// retrieve the email from the claims
	email, ok := claims["email"].(string)
	if !ok {
		c.IndentedJSON(500, gin.H{"error": "No email found in claims"})
		c.Abort()
		return
	}

	// retrieve the email from the claims
	user_id, ok := claims["user_id"].(string)
	if !ok {
		c.IndentedJSON(500, gin.H{"error": "No user id found in claims"})
		c.Abort()
		return
	}

	c.IndentedJSON(200, gin.H{"message": "HORRAH!!! You are authenticated and authorized to use this route", "email": email, "user_id": user_id})
}
