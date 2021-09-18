package infra

import (
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/go-gorp/gorp"
)

type QuestionRepository struct {
	dbmap *gorp.DbMap
}

func NewQuestionRepository(dbmap *gorp.DbMap) *QuestionRepository {
	dbmap.AddTableWithName(entity.Answer{}, "answers").SetKeys(true, "id")
	return &QuestionRepository{dbmap: dbmap}
}

func (qr *QuestionRepository) Post(question *entity.Question) error {
	if err := qr.dbmap.Insert(question); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
