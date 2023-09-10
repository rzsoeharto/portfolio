package authutils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CheckIfUserExist(c *gin.Context, username string, db *pgxpool.Pool) bool {
	var count int

	row := db.QueryRow(c, `SELECT COUNT(username) FROM users WHERE username = $1`, username)

	err := row.Scan(&count)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
