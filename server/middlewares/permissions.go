package middlewares

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func PermissionCheck(c *gin.Context) {
	var a map[string]interface{}
	db, err := database.InitDB(c)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		c.Abort()
	}

	per := c.GetString("Permission")

	fmt.Println(per)

	row := db.QueryRow(c, `SELECT permissions FROM userpermissions WHERE "referenceID" = $1`, &per)

	scanerr := row.Scan(&a)

	fmt.Println(scanerr)

	c.Next()
}
