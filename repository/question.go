package repository

import (
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
)

type Question interface {
	Post(*entity.Question) error
	FindRandomly(groupID string) (*entity.Question, error)
}
