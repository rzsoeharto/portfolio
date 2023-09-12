package middlewares

import (
	"fmt"
	"portfolio/server/database"

	"github.com/gin-gonic/gin"
)

func PermissionCheck(c *gin.Context) {
	var a map[string]interface{}
	db := database.InitDB(c)

	per := c.GetString("Permission")

	fmt.Println(per)

	row := db.QueryRow(c, `SELECT permissions FROM userpermissions WHERE "referenceID" = $1`, &per)

	scanerr := row.Scan(&a)

	fmt.Println(scanerr)

	c.Next()
}
