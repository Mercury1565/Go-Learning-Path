package data

import "Implementing_Authentication_Authorization/models"

type UserCollection struct {
	Users map[string]*models.User
}

func NewUserCollection() *UserCollection {
	return &UserCollection{
		Users: make(map[string]*models.User),
	}
}

// Global variable to store the JWT secret
var JwtSecret = []byte("your_jwt_secret")
