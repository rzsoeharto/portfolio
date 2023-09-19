package posttx

import (
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostTx(c *gin.Context, db *pgxpool.Pool, post *models.BlogPost) error {
	var sections []models.PostSection

	sections = append(sections, post.Sections...)

	tx, err := db.Begin(c)

	if err != nil {
		return err
	}

	// Inserting blog_post
	if blogErr := tx.QueryRow(c, `INSERT INTO blog_posts (title, author, published) VALUES ($1, $2, $3) RETURNING id`, post.Title, post.Author, post.Published).Scan(&post.ID); blogErr != nil {
		tx.Rollback(c)
		return blogErr
	}

	// Inserting Post_Sections
	for _, section := range sections {
		_, err := tx.Exec(c, `INSERT INTO post_sections (blog_post_id, section_type, content) VALUES ($1, $2, $3)`, post.ID, section.SectionType, section.Content)
		if err != nil {
			tx.Rollback(c)
			return err
		}
	}

	commitErr := tx.Commit(c)

	if commitErr != nil {
		tx.Rollback(c)
		return commitErr
	}

	responses.Code200(c, "New Entry Successfully Created!")

	return nil
}
