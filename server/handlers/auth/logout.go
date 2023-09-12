package auth

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		responses.Code401(c, "Missing authorization header")
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == "" {
		responses.Code401(c, "Token is invalid or expired")
		return
	}

	db := database.InitDB(c)

	defer db.Close()

	row, err := db.Exec(c, `INSERT INTO blacklist (token) VALUES ($1)`, &tokenString)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	fmt.Println(row)
	responses.Code202(c, "Logged out")
}
