package handlers

import (
	"net/http"
	"portfolio/server/database"
	"portfolio/server/jwt_utils"
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ReplenishToken(c *gin.Context, token *jwt.Token) {
	var user models.User

	db, err := database.InitDB(c)

	if err != nil {
		responses.Code500(c)
		return
	}

	defer db.Close()

	acc, refID := jwt_utils.GenerateAccessToken(c, &user)
	ref := jwt_utils.GenerateRefreshToken(c)

	db.Exec(c, `INSERT INTO userpermissions (referenceID) VALUES ($2) WHERE username = $1`, &user.Username, &refID)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":    acc,
		"referenceToken": ref,
	})
}
