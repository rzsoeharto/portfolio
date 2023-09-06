package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------------------------------------------------- 500s
func Code500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong",
	})
}

// ------------------------------------------------------- 400s
func Code404(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": msg,
	})
}

func Code401(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": msg,
	})
}

func Code400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})
}

// ------------------------------------------------------- 300s
func Code302(c *gin.Context, msg string) {
	c.JSON(http.StatusFound, gin.H{
		"message": msg,
	})
}

// ------------------------------------------------------- 200s
func Code200(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func Code202(c *gin.Context, msg string) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": msg,
	})
}
