package infra

import (
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
	"github.com/go-gorp/gorp"
)

var _ repository.Comment = (*CommentRepository)(nil)

type CommentRepository struct {
	dbmap *gorp.DbMap
}

func NewCommentRepository(dbmap *gorp.DbMap) *CommentRepository {
	dbmap.AddTableWithName(entity.Comment{}, "comments")
	return &CommentRepository{dbmap: dbmap}
}

func (c *CommentRepository) Post(comment *entity.Comment) error {
	if err := c.dbmap.Insert(comment); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (c *CommentRepository) FindUnique(groupID, questionID, answerID, commentID string) (*entity.Comment, error) {
	query := `SELECT * FROM comments
				WHERE group_id = ? AND question_id = ? AND answer_id = ? AND id = ?`

	comment := &entity.Comment{}
	if err := c.dbmap.SelectOne(comment, query, groupID, questionID, answerID, commentID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return comment, nil
}

func (c *CommentRepository) FindByAnswer(groupID, questionID, answerID string) ([]*entity.Comment, error) {
	query := `SELECT * FROM comments
				WHERE group_id = ? AND question_id = ? AND answer_id = ?
				ORDER BY created_at DESC`

	comments := []*entity.Comment{}
	if _, err := c.dbmap.Select(&comments, query, groupID, questionID, answerID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return comments, nil
}
