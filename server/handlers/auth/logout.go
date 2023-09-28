package auth

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	sid := c.GetString("Session ID")
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

	tx, err := db.Begin(c)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}
	_, txerr1 := tx.Exec(c, `INSERT INTO blacklist (token) VALUES ($1)`, &tokenString)

	if err != nil {
		fmt.Println(txerr1)
		tx.Rollback(c)
		responses.Code500(c)
		return
	}

	_, txerr2 := tx.Exec(c, `DELETE FROM sessions WHERE session_id = $1`, sid)

	if txerr2 != nil {
		tx.Rollback(c)
		fmt.Println(txerr1)
		responses.Code500(c)
		return
	}

	commitErr := tx.Commit(c)

	if commitErr != nil {
		fmt.Println(commitErr)
		responses.Code500(c)
		return
	}

	responses.Code202(c, "Logged out")
}
