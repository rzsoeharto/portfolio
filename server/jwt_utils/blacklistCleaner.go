package jwt_utils

import (
	"context"
	"fmt"
	"portfolio/server/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CleanUpTokens(c *gin.Context) {
	fmt.Println("Cleanup Initiated")
	ticker := time.NewTicker(30 * time.Minute)

	defer ticker.Stop()

	db := database.InitDB(c)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Cleaning: Underway")
			cleanExpRefTokens(db)
			fmt.Println("Cleaning: Done")
		}
	}
}

func cleanExpRefTokens(db *pgxpool.Pool) error {
	_, queryErr := db.Exec(context.Background(), `DELETE FROM blacklist WHERE blacklisted < current_timestamp - interval '1 week';`)

	if queryErr != nil {
		fmt.Println(queryErr)
		return queryErr
	}

	fmt.Println("OK")
	return nil
}
