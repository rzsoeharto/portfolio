package middlewares

import (
	"log"
	"portfolio/server/database"
	"portfolio/server/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckBlacklist(c *gin.Context) {
	var count int

	db, err := database.InitDB(c)

	if err != nil {
		log.Fatal("Connection to database failed", err)
		responses.Code500(c)
		c.Abort()
		return
	}

	defer db.Close()

	header := c.GetHeader("Authorization")

	if header == "" {
		responses.Code401(c, "Missing authorization header")
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")

	if tokenString == "" {
		responses.Code401(c, "Missing authorization header")
		c.Abort()
		return
	}

	row := db.QueryRow(c, `SELECT COUNT(token) FROM blacklist WHERE token = $1`, tokenString)

	row.Scan(&count)

	if count > 0 {
		responses.Code401(c, "Token is blacklisted")
		c.Abort()
	}

	c.Next()
}
