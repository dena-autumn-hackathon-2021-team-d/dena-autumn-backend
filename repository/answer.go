package repository

import "github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"

type Answer interface {
	Post(answer *entity.Answer) error
	FindByGroupID(groupID string) ([]*entity.Answer, error)
}
