package middlewares

import (
	"portfolio/server/database"
	logger "portfolio/server/logs"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func CheckBlacklist(c *gin.Context) {
	var count int

	db := database.InitDB(c)

	defer db.Close()

	refreshToken, cookieErr := c.Cookie("Refresh-Token")

	if cookieErr != nil {
		logger.Logger.Println("No cookie: ", cookieErr)
		responses.Code401(c, "Missing refresh cookie")
		c.Abort()
		return
	}

	row := db.QueryRow(c, `SELECT COUNT(token) FROM blacklist WHERE token = $1`, refreshToken)

	row.Scan(&count)

	if count > 0 {
		responses.Code401(c, "Token is blacklisted")
		c.Abort()
		return
	}

	c.Next()
}
