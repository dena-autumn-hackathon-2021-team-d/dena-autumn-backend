package entity

import "github.com/google/uuid"

type Answer struct {
	ID         string `json:"id" db:"id"`
	GroupID    string `json:"group_id" db:"group_id"`
	QuestionID string `json:"question_id" db:"question_id"`
	Contents   string `json:"contents" db:"contents"`
	Username   string `json:"username" db:"username"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}

func (a *Answer) NewID() {
	a.ID = uuid.New().String()
	return
}
