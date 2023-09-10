package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ReplenishTokenTx(c *gin.Context, db *pgx.Conn, sid string) error {
	var uname string
	tx, err := db.Begin(c)

	if err != nil {
		tx.Conn().Close(c)
		return err
	}

	row := tx.QueryRow(c, `SELECT username FROM session WHERE session_id = $1`, &sid)

	scanerr := row.Scan(&uname)

	if scanerr != nil {
		tx.Rollback(c)
		return err
	}

	return nil
}
