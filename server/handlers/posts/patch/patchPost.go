package patchpost

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/models"
	"portfolio/server/responses"
	posttx "portfolio/server/transactions/post"

	"github.com/gin-gonic/gin"
)

func EditPost(c *gin.Context) {
	var post models.BlogPost

	bindErr := c.BindJSON(&post)

	if bindErr != nil {
		fmt.Println(bindErr)
		responses.Code400(c, "Incomplete fields or invalid data")
		return
	}

	db := database.InitDB(c)

	err := posttx.EditPostTx(c, db, post)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	responses.Code200(c, "Changes Saved")
}
