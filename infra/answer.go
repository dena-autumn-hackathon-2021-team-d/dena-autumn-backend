package infra

import (
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/go-gorp/gorp"
)

type AnswerRepository struct {
	dbmap *gorp.DbMap
}

func NewAnswerRepository(dbmap *gorp.DbMap) *AnswerRepository {
	dbmap.AddTableWithName(entity.Answer{}, "answers").SetKeys(true, "id", "group_id", "question_id")
	return &AnswerRepository{dbmap: dbmap}
}

func (a *AnswerRepository) Post(answer *entity.Answer) error {
	if err := a.dbmap.Insert(answer); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
