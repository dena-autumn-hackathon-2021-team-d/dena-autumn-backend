package infra

import (
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/go-gorp/gorp"
)

type CommentRepository struct {
	dbmap *gorp.DbMap
}

func NewCommentRepository(dbmap *gorp.DbMap) *CommentRepository {
	dbmap.AddTableWithName(entity.Comment{}, "comments").SetKeys(true, "id")
	return &CommentRepository{dbmap: dbmap}
}

func (c *CommentRepository) Post(comment *entity.Comment) error {
	if err := c.dbmap.Insert(comment); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
