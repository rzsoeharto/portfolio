package transactions

import (
	"portfolio/server/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(c *gin.Context, tx pgx.Tx, dbUser *models.User, user *models.User) error {
	row := tx.QueryRow(c, `SELECT password from "users" WHERE username = $1`, &user.Username)

	scanErr := row.Scan(&dbUser.Password)

	if scanErr != nil {

		return scanErr
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if passErr != nil {

		return passErr
	}

	return nil
}

func FetchUsername(c *gin.Context, tx pgx.Tx, user *models.User) error {
	data := tx.QueryRow(c, `SELECT name FROM "users" where username = $1`, &user.Username)

	scanErr := data.Scan(&user.Name)

	if scanErr != nil {

		return scanErr
	}

	return nil
}

func PermissionAndSession(c *gin.Context, tx pgx.Tx, sid string, refID string, user *models.User) error {

	_, perErr := tx.Exec(c, `UPDATE userpermissions SET "referenceID" = $1, "active" = $2 WHERE "username" = $3;`, &refID, true, &user.Username)

	if perErr != nil {

		return perErr
	}

	_, sesErr := tx.Exec(c, `INSERT INTO sessions (session_id, username) VALUES($1, $2)`, &sid, &user.Username)

	if sesErr != nil {

		return sesErr
	}

	return nil
}
