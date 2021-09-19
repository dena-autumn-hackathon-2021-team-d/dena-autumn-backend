package entity

import "github.com/google/uuid"

type Comment struct {
	ID         string `json:"id" db:"id"`
	GroupID    string `json:"group_id" db:"group_id"`
	QuestionID string  `json:"question_id" db:"question_id"`
	AnswerID   string `json:"answer_id" db:"answer_id"`
	Contents   string `json:"contents" db:"contents"`
	Username   string `json:"username" db:"username"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}

func (c *Comment) NewID() {
	c.ID = uuid.New().String()
	return
}
