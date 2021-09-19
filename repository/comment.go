package repository

import "github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"

type Comment interface {
	Post(comment *entity.Comment) error
	FindUnique(groupID, questionID, answerID, commentID string) (*entity.Comment, error)
	FindByAnswer(groupID, questionID, answerID string) ([]*entity.Comment, error)
}
