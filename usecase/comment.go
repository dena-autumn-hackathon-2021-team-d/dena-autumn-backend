package usecase

import (
	"fmt"
	"time"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
)

type CommentUseCase struct {
	commentRepo repository.Comment
}

func NewCommentUseCase(commentRepo repository.Comment) *CommentUseCase {
	return &CommentUseCase{commentRepo: commentRepo}
}

func (c *CommentUseCase) Post(comment *entity.Comment) (*entity.Comment, error) {
	comment.CreatedAt = time.Now().Format(time.RFC3339)
	comment.NewID()

	if err := c.commentRepo.Post(comment); err != nil {
		return nil, fmt.Errorf("failed to post comment: %w", err)
	}

	return comment, nil
}

func (c *CommentUseCase) GetUnique(groupID, questionID, answerID, commentID string) (*entity.Comment, error) {
	return c.commentRepo.FindUnique(groupID, questionID, answerID, commentID)
}

func (c *CommentUseCase) GetByAnswer(groupID, questionID, answerID string) ([]*entity.Comment, error) {
	return c.commentRepo.FindByAnswer(groupID, questionID, answerID)
}
