package infra

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
	"github.com/go-gorp/gorp"
)

var _ repository.Question = (*QuestionRepository)(nil)

type QuestionRepository struct {
	dbmap *gorp.DbMap
}

// QuestionDTO はNumAnswersをInsertでは無視して，Findでは利用できるようにするための構造体
type QuestionDTO struct {
	ID         string `json:"id" db:"id"`
	Contents   string `json:"contents" db:"contents"`
	GroupID    string `json:"group_id" db:"group_id"`
	Username   string `json:"username" db:"username"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	NumAnswers int    `json:"num_answers" db:"num_answers"`
}

func NewQuestionRepository(dbmap *gorp.DbMap) *QuestionRepository {
	dbmap.AddTableWithName(entity.Question{}, "questions")
	return &QuestionRepository{dbmap: dbmap}
}

func (qr *QuestionRepository) Post(question *entity.Question) error {
	if err := qr.dbmap.Insert(question); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (qr *QuestionRepository) FindRandomly(groupID string) (*entity.Question, error) {
	query := `SELECT id, contents, group_id, username, created_at, (SELECT COUNT(id) FROM answers AS a WHERE a.question_id = q.id) as num_answers
				FROM questions AS q
				WHERE group_id = ? ORDER BY RANDOM() LIMIT 1`

	question := &QuestionDTO{}
	if err := qr.dbmap.SelectOne(question, query, groupID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrQuestionNotFound
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &entity.Question{
		ID:         question.ID,
		Contents:   question.Contents,
		GroupID:    question.GroupID,
		Username:   question.Username,
		CreatedAt:  question.CreatedAt,
		NumAnswers: question.NumAnswers,
	}, nil
}

func (qr QuestionRepository) FindByQuestion(groupID, questionID string) (*entity.Question, error) {
	query := `SELECT id, contents, group_id, username, created_at, (SELECT COUNT(id) FROM answers AS a WHERE a.question_id = q.id) as num_answers
				FROM questions AS q
				WHERE id = ? AND group_id = ?`

	question := &QuestionDTO{}
	if err := qr.dbmap.SelectOne(question, query, questionID, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &entity.Question{
		ID:         question.ID,
		Contents:   question.Contents,
		GroupID:    question.GroupID,
		Username:   question.Username,
		CreatedAt:  question.CreatedAt,
		NumAnswers: question.NumAnswers,
	}, nil
}

func (qr QuestionRepository) GetAll(groupID string) ([]*entity.Question, error) {
	query := `SELECT id, contents, group_id, username, created_at, (SELECT COUNT(id) FROM answers AS a WHERE a.question_id = q.id) as num_answers
				FROM questions AS q
				WHERE group_id = ?
				ORDER BY created_at DESC`

	questions := []*QuestionDTO{}
	if _, err := qr.dbmap.Select(&questions, query, groupID); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	resQuestions := make([]*entity.Question, len(questions))
	for i, question := range questions {
		resQuestions[i] = &entity.Question{
			ID:         question.ID,
			Contents:   question.Contents,
			GroupID:    question.GroupID,
			Username:   question.Username,
			CreatedAt:  question.CreatedAt,
			NumAnswers: question.NumAnswers,
		}
	}
	return resQuestions, nil
}
