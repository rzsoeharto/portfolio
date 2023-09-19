package posttx

import (
	"portfolio/server/models"
	"portfolio/server/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func EditPostTx(c *gin.Context, db *pgxpool.Pool, post models.BlogPost) error {
	var sections []models.PostSection

	sections = append(sections, post.Sections...)

	tx, err := db.Begin(c)

	if err != nil {
		return err
	}

	tNow := time.Now()

	_, editErr := tx.Exec(c, `UPDATE blog_posts SET title = $1, updated = $2 WHERE "id" = $3`, post.Title, tNow, post.ID)

	if editErr != nil {
		tx.Rollback(c)
		responses.Code500Message(c, "Something went wrong. Changes discarded")
		return editErr
	}

	for _, section := range sections {
		_, err := tx.Exec(c, `UPDATE post_sections SET section_type = $1, content = $2 WHERE id = $3`, section.SectionType, section.Content, section.ID)
		if err != nil {
			tx.Rollback(c)
			responses.Code500Message(c, "Something went wrong. Changes discarded")
			return err
		}
	}

	commitErr := tx.Commit(c)

	if commitErr != nil {
		tx.Rollback(c)
		responses.Code500Message(c, "Something went wrong. Changes discarded")
		return commitErr
	}

	return nil
}
