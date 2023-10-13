package handlers

import (
	"fmt"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/jwt_utils"
	"portfolio/server/responses"
	transactions "portfolio/server/transactions/auth"

	"github.com/gin-gonic/gin"
)

func ReplenishToken(c *gin.Context) {
	db := database.InitDB(c)

	defer db.Close()

	tokenSID := c.GetString("Session ID")

	acc, refID := jwt_utils.GenerateAccessToken(c)
	ref, newSID := jwt_utils.GenerateRefreshToken(c)

	txerr := transactions.ReplenishTokenTx(c, db, tokenSID, refID, newSID)

	if txerr != nil {
		fmt.Println(txerr)
		responses.Code500(c)
		return
	}

	c.Header("Authorization", "Bearer "+acc)

	c.Header("Refresh-Token", "Refresh "+ref)

	c.JSON(http.StatusOK, gin.H{
		"Message": "All good",
	})
}
