package auth

import (
	"portfolio/server/database"
	logger "portfolio/server/logs"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	sid := c.GetString("Session ID")
	refreshToken, cookieErr := c.Cookie("Refresh-Token")

	if cookieErr != nil {
		responses.Code401(c, "Missing refresh cookie")
		c.Abort()
		return
	}

	if refreshToken == "" {
		responses.Code401(c, "Missing authorization token")
		c.Abort()
		return
	}

	db := database.InitDB(c)

	defer db.Close()

	tx, err := db.Begin(c)

	if err != nil {
		logger.Logger.Fatalln(err)
		responses.Code500(c)
		return
	}
	_, txerr1 := tx.Exec(c, `INSERT INTO blacklist (token) VALUES ($1)`, refreshToken)

	if err != nil {
		logger.Logger.Fatalln(txerr1)
		tx.Rollback(c)
		responses.Code500(c)
		return
	}

	_, txerr2 := tx.Exec(c, `DELETE FROM sessions WHERE session_id = $1`, sid)

	if txerr2 != nil {
		tx.Rollback(c)
		logger.Logger.Fatalln(txerr1)
		responses.Code500(c)
		return
	}

	commitErr := tx.Commit(c)

	if commitErr != nil {
		logger.Logger.Fatalln(commitErr)
		responses.Code500(c)
		tx.Rollback(c)
		return
	}

	c.SetSameSite(4)
	c.SetCookie("Authorization", "", -1, "/", "localhost", true, true)
	c.SetCookie("Refresh-Token", "", -1, "/", "localhost", true, true)

	responses.Code202(c, "Logged out")
}
