package auth

import (
	"fmt"
	"log"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/jwt_utils"
	"portfolio/server/models"
	"portfolio/server/responses"
	"portfolio/server/utils"
	authutils "portfolio/server/utils/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var dbUser models.User

	db, err := database.InitDB(c)

	if err != nil {
		log.Fatal("Connection to database failed: ", err)
		responses.Code500(c)
		return
	}

	defer db.Close()

	if !utils.BindJSON(c, &user) {
		return
	}

	if !authutils.CheckIfUserExist(c, user.Username, db) {
		responses.Code404(c, "User does not exist")
		return
	}

	row := db.QueryRow(c, `SELECT password FROM "users" WHERE username = $1`, &user.Username)

	scanerr := row.Scan(&dbUser.Password)

	if scanerr != nil {
		fmt.Println(scanerr)
		responses.Code500(c)
		return
	}

	passerr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if passerr != nil {
		responses.Code400(c, "wrong password")
		return
	}

	acc, refID := jwt_utils.GenerateAccessToken(c, user.Username)
	ref, sid := jwt_utils.GenerateRefreshToken(c)

	data := db.QueryRow(c, `SELECT name FROM "users" where username = $1`, &user.Username)

	scanErr := data.Scan(&user.Name)

	if scanErr != nil {
		fmt.Println(scanErr)
		responses.Code500(c)
		return
	}

	per, err := db.Exec(c, `UPDATE userpermissions
	SET "referenceID" = $1, "active" = $2
	WHERE "username" = $3;`, &refID, true, &user.Username)

	fmt.Println(per)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	ses, err := db.Exec(c, `INSERT INTO sessions (session_id, username) VALUES($1, $2)`, &sid, &user.Username)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	fmt.Println(ses)

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"name":     user.Name,
		"access":   acc,
		"refresh":  ref,
	})
}
