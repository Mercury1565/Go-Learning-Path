package router

import (
	"Implementing_Authentication_Authorization/controller"
	"Implementing_Authentication_Authorization/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controller.HandelUserRegister)
	r.POST("/login", controller.HandelUserLogin)

	r.GET("/secure", middleware.AuthMiddleware(), controller.HandelSecureMessage)

	return r

}
