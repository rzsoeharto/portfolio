package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CheckIfUserExist(c *gin.Context, username string, db *pgxpool.Pool) bool {
	var count int

	row := db.QueryRow(c, `SELECT username FROM users WHERE username = $1`, username)

	err := row.Scan(&count)

	if errors.Is(err, pgx.ErrNoRows) {
		return false
	}

	return true
}
