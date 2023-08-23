package auth

import (
	"portfolio/server/models"
	"portfolio/server/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	// var dbUser models.User

	if !utils.BindJSON(c, &user) {
		return
	}

}
