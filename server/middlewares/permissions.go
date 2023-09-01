package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PermissionCheck(c *gin.Context) {
	fmt.Println("Hi")
	c.Next()
}
