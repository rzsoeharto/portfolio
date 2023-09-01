package utils

import (
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	err := c.BindJSON(data)
	if err != nil {
		responses.Code400(c, "Invalid data or incomplete fields")
		return false
	}
	return true
}
