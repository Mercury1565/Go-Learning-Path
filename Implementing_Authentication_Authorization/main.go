package main

import "Implementing_Authentication_Authorization/router"

func main() {

	route := router.SetUpRouter()
	route.Run("localhost:8080")
}
