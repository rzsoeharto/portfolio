package delpost

import (
	"fmt"
	"portfolio/server/database"
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	var post models.BlogPost

	bindErr := c.BindJSON(&post)

	if bindErr != nil {
		fmt.Println(bindErr)
		responses.Code400(c, "Invalid data or incomplete fields")
		return
	}

	db := database.InitDB(c)

	tx, err := db.Begin(c)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	_, txErr := tx.Exec(c, `DELETE FROM blog_posts WHERE id = $1`, post.ID)

	if txErr != nil {
		fmt.Println(txErr)
		tx.Rollback(c)
		responses.Code500Message(c, "Something went wrong. Changes Discarded")
		return
	}

	_, txErr2 := tx.Exec(c, `DELETE FROM post_sections WHERE blog_post_id = $1`, post.ID)

	if txErr2 != nil {
		fmt.Println(txErr2)
		tx.Rollback(c)
		responses.Code500Message(c, "Something went wrong. Changes Discarded")
		return
	}

	comErr := tx.Commit(c)

	if comErr != nil {
		fmt.Println(comErr)
		tx.Rollback(c)
		responses.Code500Message(c, "Something went wrong. Changes Discarded")
		return
	}

	responses.Code200(c, "Post Successfully Deleted")
}
