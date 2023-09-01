package database

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(c *gin.Context) (*pgxpool.Pool, error) {
	conf, parseErr := pgxpool.ParseConfig(os.Getenv("DB_URL"))

	if parseErr != nil {
		return nil, parseErr
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)

	if err != nil {
		return nil, err
	}

	return pool, nil

}
