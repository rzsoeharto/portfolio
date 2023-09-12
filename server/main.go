package main

import (
	"portfolio/server/handlers/auth"
	handlers "portfolio/server/handlers/jwt_handlers"
	"portfolio/server/initializers"
	"portfolio/server/jwt_utils"
	"portfolio/server/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	go jwt_utils.CleanUpTokens(&gin.Context{})

	// GET Endpoints

	//POST Endpoints
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.POST("/logout", middlewares.RefreshValidator, auth.Logout)
	r.POST("/Replenish", middlewares.CheckBlacklist, middlewares.RefreshValidator, handlers.ReplenishToken)
	// PATCH Endpoints

	// DELETE Endpoints

	r.Run()
}
