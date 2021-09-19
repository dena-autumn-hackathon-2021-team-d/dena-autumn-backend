package repository

import "github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"

type Answer interface {
	Post(answer *entity.Answer) error
	FindByGroupID(groupID string) ([]*entity.Answer, error)
	FindUnique(groupID, questionID, answerID string) (*entity.Answer, error)
	FindByQuestion(groupID, questionID string) ([]*entity.Answer, error)
}
