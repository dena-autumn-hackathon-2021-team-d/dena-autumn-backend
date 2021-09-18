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
	dbmap.AddTableWithName(entity.Question{}, "questions").SetKeys(true, "id")
	return &QuestionRepository{dbmap: dbmap}
}

func (qr *QuestionRepository) Post(question *entity.Question) error {
	if err := qr.dbmap.Insert(question); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (qr *QuestionRepository) FindRandomly(groupID string) (*entity.Question, error) {
	query := `SELECT * FROM questions WHERE group_id = ? ORDER BY RAND() LIMIT 1`

	var question *entity.Question
	if err := qr.dbmap.SelectOne(&question, query, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return question, nil
}

func (qr QuestionRepository) FindByQuestion(groupID string, questionID int) (*entity.Question, error) {
	query := `SELECT * FROM questions WHERE question_id = ? AND group_id = ?`

	var questions *entity.Question
	if err := qr.dbmap.SelectOne(questions, query, questionID, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return questions, nil
}

func (qr QuestionRepository) GetAll(groupID string) ([]*entity.Question, error) {
	query := `SELECT * FROM questions WHERE group_id = ? ORDER BY created_at DESC `

	var questions []*entity.Question
	if _, err := qr.dbmap.Select(&questions, query, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return questions, nil
}
