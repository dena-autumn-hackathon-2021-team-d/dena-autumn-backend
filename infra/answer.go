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
	dbmap.AddTableWithName(entity.Answer{}, "answers").SetKeys(true, "id")
	return &AnswerRepository{dbmap: dbmap}
}

func (a *AnswerRepository) Post(answer *entity.Answer) error {
	if err := a.dbmap.Insert(answer); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (a *AnswerRepository) FindByGroupID(groupID string) ([]*entity.Answer, error) {
	query := `SELECT * FROM answers WHERE group_id = ?`

	answers := []*entity.Answer{}
	if _, err := a.dbmap.Select(&answers, query, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return answers, nil
}

func (a *AnswerRepository) FindUnique(groupID string, questionID, answerID int) (*entity.Answer, error) {
	query := `SELECT * FROM answers WHERE group_id = ? AND question_id = ? AND id = ?`

	answer := &entity.Answer{}
	if err := a.dbmap.SelectOne(answer, query, groupID, questionID, answerID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return answer, nil
}

func (a *AnswerRepository) FindByQuestion(groupID string, questionID int) ([]*entity.Answer, error) {
	query := `SELECT * FROM answers WHERE group_id = ? AND question_id = ?`

	answers := []*entity.Answer{}
	if _, err := a.dbmap.Select(&answers, query, groupID, questionID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return answers, nil
}
