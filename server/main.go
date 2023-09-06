package main

import (
	"fmt"
	"portfolio/server/handlers/auth"
	"portfolio/server/initializers"
	"portfolio/server/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func test(c *gin.Context) {
	per := c.GetString("Permission")
	fmt.Println(per)
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func main() {
	r := gin.Default()

	// GET Endpoints

	//POST Endpoints
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.POST("/logout", middlewares.RefreshValidator, auth.Logout)
	r.POST("/", middlewares.PermissionCheck, middlewares.AccessValidator, test)
	// PATCH Endpoints

	// DELETE Endpoints

	r.Run()
}
