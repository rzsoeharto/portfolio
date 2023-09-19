package transactions

import (
	"portfolio/server/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoginTx(c *gin.Context, db *pgxpool.Pool, refid string, sid string, user models.User) error {
	tx, _ := db.Begin(c)

	_, perErr := tx.Exec(c, `UPDATE userpermissions
	SET "referenceID" = $1, "active" = $2
	WHERE "username" = $3;`, refid, true, user.Username)

	if perErr != nil {
		tx.Rollback(c)
		return perErr
	}

	_, sesErr := tx.Exec(c, `INSERT INTO sessions (session_id, username) VALUES($1, $2)`, &sid, &user.Username)

	if sesErr != nil {
		tx.Rollback(c)
		return perErr
	}

	return nil
}
