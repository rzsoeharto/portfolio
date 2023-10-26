package middlewares

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func containString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

func PermissionCheck(c *gin.Context) {
	var a []string

	db := database.InitDB(c)

	per := c.GetString("Permission")

	row := db.QueryRow(c, `SELECT permissions FROM userpermissions WHERE "referenceID" = $1`, &per)

	scanerr := row.Scan(&a)

	if scanerr != nil {
		fmt.Println(scanerr)
		responses.Code500Message(c, "Database error")
		c.Abort()
		return
	}

	target := "post"

	found := containString(a, target)

	if !found {
		fmt.Println("not found")
		responses.Code401(c, "Unauthorised")
		c.Abort()
		return
	}

	c.Next()
}
