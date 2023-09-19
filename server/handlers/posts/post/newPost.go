package newpost

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/models"
	"portfolio/server/responses"
	posttx "portfolio/server/transactions/post"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.BlogPost

	bindErr := c.BindJSON(&post)

	if bindErr != nil {
		fmt.Println(bindErr)
		responses.Code400(c, "Incomplete fields or invalid data")
		return
	}

	db := database.InitDB(c)

	post.Published = time.Now()
	post.Author = "Rizky"

	err := posttx.NewPostTx(c, db, &post)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		db.Close()
		return
	}

	defer db.Close()
}
