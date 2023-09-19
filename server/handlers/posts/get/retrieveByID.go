package getposts

import (
	"fmt"
	"net/http"
	"portfolio/server/database"
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
)

func RetrievePostByID(c *gin.Context) {
	var blogPost models.BlogPost

	db := database.InitDB(c)

	id := c.Param("id")

	err := db.QueryRow(c, `SELECT id, author ,title, published FROM blog_posts WHERE id = $1`, id).Scan(&blogPost.ID, &blogPost.Author, &blogPost.Title, &blogPost.Published)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	rows, err := db.Query(c, `SELECT id, blog_post_id, section_type, content FROM post_sections WHERE blog_post_id = $1`, id)

	if err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	for rows.Next() {
		var section models.PostSection

		if err := rows.Scan(&section.ID, &section.BlogPostID, &section.SectionType, &section.Content); err != nil {
			fmt.Println(err)
			return
		}
		blogPost.Sections = append(blogPost.Sections, section)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		responses.Code500(c)
		return
	}

	c.JSON(http.StatusOK, blogPost)
}
