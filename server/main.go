package main

import (
	"portfolio/server/handlers/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET Endpoints

	//POST Endpoints
	r.POST("/login", auth.Login)

	// PATCH Endpoints

	// DELETE Endpoints

	r.Run()
}
