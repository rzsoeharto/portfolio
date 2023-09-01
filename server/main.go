package main

import (
	"portfolio/server/handlers/auth"
	"portfolio/server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	// GET Endpoints

	//POST Endpoints
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	// PATCH Endpoints

	// DELETE Endpoints

	r.Run()
}
