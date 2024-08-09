package controller

import (
	"Implementing_Authentication_Authorization/data"
	"Implementing_Authentication_Authorization/models"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *data.UserCollection

func init() {
	UserCollection = data.NewUserCollection()
}

func BindJSON(user *models.User, context *gin.Context) error {
	err := context.ShouldBindJSON(&user)
	return err
}

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashedPassword, err
}

func ValidatePassword(curr_password string, existing_password string) error {
	return bcrypt.CompareHashAndPassword([]byte(existing_password), []byte(curr_password))
}

func GenerateSignedToken(user_id string, user_email string) (string, error) {
	// Generate JWT with claims of 'user_id' and 'user_email'
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email":   user_email,
	})

	// sign the token with the secret key 'jwtSecret'
	return token.SignedString(data.JwtSecret) // [JWT token]string, [err]error
}

func GetClaimsFromContext(context *gin.Context) (string, string, error) {
	// retrieve claims from the context
	claimsValue, exists := context.Get("claims")

	if !exists {
		return "", "", errors.New("no claims found")
	}

	// retrieve jwt.MapClaims from the claimsValue
	claims, ok := claimsValue.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("claims are not valid")
	}

	// retrieve the email from the claims
	email, ok := claims["email"].(string)
	if !ok {
		return "", "", errors.New("no email found in claims")
	}

	// retrieve the email from the claims
	user_id, ok := claims["user_id"].(string)
	if !ok {
		return "", "", errors.New("no user_id found in claims")
	}

	return user_id, email, nil
}

func HandelUserRegister(context *gin.Context) {
	var curr_user models.User

	// get inputed info from body
	err := BindJSON(&curr_user, context)

	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "invalid request"})
		return
	}

	// hash inputed password
	hashedPassword, err := HashPassword(curr_user.Password)

	if err != nil {
		context.IndentedJSON(500, gin.H{"error": "internal server error"})
		return
	}

	// store info for the newly registered user
	curr_user.Password = string(hashedPassword)
	UserCollection.Users[curr_user.Email] = &curr_user

	context.IndentedJSON(200, gin.H{"message": "user registered successfully"})
}

func HandelUserLogin(context *gin.Context) {
	var curr_user models.User
	// get inputed info from body
	err := context.ShouldBindJSON(&curr_user)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "invalid request"})
		return
	}

	// fetch user
	existingUser, ok := UserCollection.Users[curr_user.Email]
	if !ok {
		context.IndentedJSON(401, gin.H{"error": "email doesn't exist"})
		return
	}

	// check if user has inputed the correct password
	err = ValidatePassword(curr_user.Password, existingUser.Password)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "Incorrect password"})
		return
	}

	// generate signed JWT with 'user_id' and 'user_email' claims
	signed_jwt_token, err := GenerateSignedToken(existingUser.ID, existingUser.Email)
	if err != nil {
		context.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	context.IndentedJSON(200, gin.H{"message": "user logged in successfully", "token": signed_jwt_token})
}

func HandelSecureMessage(context *gin.Context) {
	// get the claims used to generate the token
	user_id, email, err := GetClaimsFromContext(context)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
	}

	context.IndentedJSON(200, gin.H{"message": "YES!!! You are authenticated and authorized to use this route", "email": email, "user_id": user_id})
}
