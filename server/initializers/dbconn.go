package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func InitDB(c *gin.Context) *pgx.Conn {
	db, err := pgx.Connect(c, os.Getenv(key))
	return db
}
