package database

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(c *gin.Context) *pgxpool.Pool {
	conf, parseErr := pgxpool.ParseConfig(os.Getenv("DB_URL"))

	if parseErr != nil {
		log.Panic(parseErr)
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)

	if err != nil {
		log.Panic(err)
		return nil
	}

	return pool

}
