package main

import (
	"portfolio/server/handlers/auth"
	handlers "portfolio/server/handlers/jwt_handlers"
	delpost "portfolio/server/handlers/posts/delete"
	getposts "portfolio/server/handlers/posts/get"
	patchpost "portfolio/server/handlers/posts/patch"
	newpost "portfolio/server/handlers/posts/post"
	"portfolio/server/initializers"
	"portfolio/server/jwt_utils"
	"portfolio/server/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	go jwt_utils.CleanUpTokens(&gin.Context{})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// GET Endpoints
	r.GET("/post/:id", getposts.RetrievePostByID)
	r.GET("/posts", getposts.RetrieveAll)

	//POST Endpoints
	// -------------------------------------------- Auth endpoints ----------------------------------------
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.POST("/logout", middlewares.RefreshValidator, auth.Logout)
	r.POST("/replenish", middlewares.CheckBlacklist, middlewares.RefreshValidator, handlers.ReplenishToken)

	// ------------------------------------------ Post endpoint(s) ----------------------------------------
	r.POST("/create-post", middlewares.AccessValidator, middlewares.PermissionCheck, newpost.CreatePost)

	// PATCH Endpoints
	r.PATCH("/edit-post", middlewares.AccessValidator, middlewares.PermissionCheck, patchpost.EditPost)

	// DELETE Endpoints
	r.DELETE("/delete-post", middlewares.AccessValidator, middlewares.PermissionCheck, delpost.DeletePost)

	// Run
	r.Run()
}
