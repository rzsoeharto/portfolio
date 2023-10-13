package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ReplenishTokenTx(c *gin.Context, db *pgxpool.Pool, tokenSID string, refid string, newSID string) error {
	var uname string

	refToken := c.GetString("Refresh Token")

	tx, err := db.Begin(c)

	if err != nil {
		return err
	}

	row := tx.QueryRow(c, `SELECT username FROM sessions WHERE session_id = $1`, tokenSID)

	scanErr := row.Scan(&uname)

	if scanErr != nil {
		tx.Rollback(c)
		return scanErr
	}

	_, err = tx.Exec(c, `UPDATE userpermissions SET "referenceID" = $2 WHERE username = $1`, uname, refid)

	if err != nil {
		tx.Rollback(c)
		return err
	}

	_, err = tx.Exec(c, `UPDATE sessions set session_id = $1 WHERE username = $2`, newSID, uname)

	if err != nil {
		tx.Rollback(c)
		return err
	}

	// Double check before deployment
	_, txErr1 := tx.Exec(c, `INSERT INTO blacklist (token) VALUES ($1)`, refToken)

	if txErr1 != nil {
		tx.Rollback(c)
		return txErr1
	}

	txCommitErr := tx.Commit(c)

	if txCommitErr != nil {
		tx.Rollback(c)
		return txCommitErr
	}

	return nil
}
