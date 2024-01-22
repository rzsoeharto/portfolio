package auth

import (
	"fmt"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/jwt_utils"
	logger "portfolio/server/logs"
	"portfolio/server/models"
	"portfolio/server/responses"
	transactions "portfolio/server/transactions/auth"
	"portfolio/server/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var dbUser models.User

	db := database.InitDB(c)

	defer db.Close()

	bindErr := c.BindJSON(&user)

	if bindErr != nil {
		logger.Logger.Fatalln(bindErr)
		responses.Code400(c, "Incomplete fields or invalid data")
		return
	}

	if !utils.CheckIfUserExist(c, user.Username, db) {
		responses.Code404(c, "User does not exist")
		return
	}

	tx, err := db.Begin(c)

	if err != nil {
		fmt.Println(err)
		logger.Logger.Fatalln("Error starting transaction:", err)
		responses.Code500(c)
		return
	}

	passErr := transactions.ValidatePassword(c, tx, &dbUser, &user)

	if passErr != nil {
		responses.Code401(c, "Incorrect username or password")
		logger.Logger.Println("Error in password validation", passErr)
		tx.Rollback(c)
		return
	}

	acc, refID := jwt_utils.GenerateAccessToken(c)
	ref, sid := jwt_utils.GenerateRefreshToken(c)

	usernameErr := transactions.FetchUsername(c, tx, &user)

	if usernameErr != nil {
		logger.Logger.Println("Error fetching username", usernameErr)
		responses.Code500(c)
		return
	}

	perErr := transactions.PermissionAndSession(c, tx, sid, refID, &user)

	if perErr != nil {
		responses.Code500(c)
		logger.Logger.Println("Error updating permisions and session", perErr)
		tx.Rollback(c)
		return
	}

	commitErr := tx.Commit(c)

	if commitErr != nil {
		logger.Logger.Fatal("Error committing: ", commitErr)
		responses.Code500(c)
		tx.Rollback(c)
		return
	}

	c.SetSameSite(4)
	c.SetCookie("Authorization", acc, 3600, "/", "localhost", true, true)
	c.SetCookie("Refresh-Token", ref, 604800, "/", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"name":     user.Name,
	})
}
