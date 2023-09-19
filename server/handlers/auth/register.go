package auth

import (
	"fmt"
	"portfolio/server/database"
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
		fmt.Println(bindErr)
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
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	status, err := db.Exec(c, `INSERT INTO "users" (username, password) VALUES ($1, $2)`, &user.Username, &hashedPassword)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	fmt.Println(status)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	responses.Code200(c, "Account successfully created")
}
