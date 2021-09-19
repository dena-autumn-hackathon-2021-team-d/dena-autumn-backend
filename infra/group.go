package infra

import (
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
	"github.com/go-gorp/gorp"
)

var _ repository.Group = (*GroupRepository)(nil)

type GroupRepository struct {
	dbMap *gorp.DbMap
}

func NewGroupRepository(dbMap *gorp.DbMap) *GroupRepository {
	dbMap.AddTableWithName(entity.Group{}, "groups")
	return &GroupRepository{
		dbMap: dbMap,
	}
}

func (r *GroupRepository) Insert(group *entity.Group) error {
	if err := r.dbMap.Insert(group); err != nil {
		fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *GroupRepository) FindByID(groupID string) (*entity.Group, error) {
	query := `SELECT * FROM groups WHERE id = ?`

	var group *entity.Group
	if err := r.dbMap.SelectOne(group, query, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return group, nil
}
