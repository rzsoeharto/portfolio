package getposts

import (
	"fmt"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func RetrieveAll(c *gin.Context) {
	var postSlice []models.BlogPost

	db := database.InitDB(c)

	defer db.Close()

	rows, err := db.Query(c, `SELECT id, title, published FROM blog_posts`)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	for rows.Next() {
		var post models.BlogPost

		if err != nil {
			fmt.Println(err)
			responses.Code500(c)
			return
		}

		if err := rows.Scan(&post.ID, &post.Title, &post.Published); err != nil {
			fmt.Println(err)
			return
		}

		postSlice = append(postSlice, post)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	c.JSON(http.StatusOK, postSlice)
}
