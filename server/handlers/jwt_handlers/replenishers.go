package handlers

import (
	"log"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/jwt_utils"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func ReplenishToken(c *gin.Context) {
	var uname string
	db, err := database.InitDB(c)

	if err != nil {
		responses.Code500(c)
		return
	}

	defer db.Close()

	tokenSID := c.GetString("Session ID")

	acc, refID := jwt_utils.GenerateAccessToken(c)

	ref, newSID := jwt_utils.GenerateRefreshToken(c)

	row := db.QueryRow(c, `SELECT username FROM sessions WHERE session_id = $1`, &newSID)

	scanErr := row.Scan(&uname)

	if scanErr != nil {
		log.Fatal(scanErr)
		c.JSON()
	}

	db.Exec(c, `INSERT INTO userpermissions (referenceID) VALUES ($2) WHERE username = $1`, &uname, &refID)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":    acc,
		"referenceToken": ref,
	})

}
