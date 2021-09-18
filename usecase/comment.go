package usecase

import (
	"fmt"
	"strconv"
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

	if err := c.commentRepo.Post(comment); err != nil {
		return nil, fmt.Errorf("failed to post comment: %w", err)
	}

	return comment, nil
}

func (c *CommentUseCase) GetUnique(groupID, questionID, answerID, commentID string) (*entity.Comment, error) {
	iQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse question id: %w", err)
	}

	iAnswerID, err := strconv.Atoi(answerID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse answer id: %w", err)
	}

	iCommentID, err := strconv.Atoi(commentID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse comment id: %w", err)
	}

	return c.commentRepo.FindUnique(groupID, iQuestionID, iAnswerID, iCommentID)
}
