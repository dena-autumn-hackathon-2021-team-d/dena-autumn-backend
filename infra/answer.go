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
	return &AnswerRepository{dbmap: dbmap}
}

func (a *AnswerRepository) Post(answer *entity.Answer) error {
	query := `INSERT INTO answers (group_id, question_id, contents, username, created_at)
				VALUE (?, ?, ?, ?, ?)`

	res, err := a.dbmap.Exec(
		query,
		answer.GroupID,
		answer.QuestionID,
		answer.Contents,
		answer.Username,
		answer.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	answer.ID = int(id)

	return nil
}
