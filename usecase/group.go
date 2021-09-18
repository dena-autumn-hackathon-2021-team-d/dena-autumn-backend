package usecase

import (
	"fmt"
	"time"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
)

type GroupUseCase struct {
	group repository.Group
}

func NewGroupUseCase(group repository.Group) *GroupUseCase {
	return &GroupUseCase{
		group: group,
	}
}

func (u *GroupUseCase) Create(group *entity.Group) error {
	group.CreatedAt = time.Now().Format(time.RFC3339)
	group.NewID()

	if err := u.group.Insert(group); err != nil {
		return fmt.Errorf("failed to Insert Group into DB: %w", err)
	}
	return nil
}
