package auth

import (
	"portfolio/server/database"
	logger "portfolio/server/logs"
	"portfolio/server/models"
	"portfolio/server/responses"
	"portfolio/server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	db := database.InitDB(c)

	defer db.Close()

	bindErr := c.BindJSON(&user)

	if bindErr != nil {
		logger.Logger.Println(bindErr)
		responses.Code400(c, "Incomplete fields or invalid data")
		return
	}

	if utils.CheckIfUserExist(c, user.Username, db) {
		responses.Code302(c, "Username is taken")
		return
	}

	if strings.Contains(user.Password, " ") {
		responses.Code400(c, "Please do not use spaces")
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		logger.Logger.Println(err)
		responses.Code500(c)
		return
	}

	_, execErr := db.Exec(c, `INSERT INTO "users" (username, password) VALUES ($1, $2)`, &user.Username, &hashedPassword)

	if execErr != nil {
		logger.Logger.Println(err)
		responses.Code500(c)
		return
	}

	if err != nil {
		logger.Logger.Println(err)
		responses.Code500(c)
		return
	}

	responses.Code200(c, "Account successfully created")
}
