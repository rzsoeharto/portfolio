package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data or incomplete fields",
		})
		return false
	}
	return true
}
