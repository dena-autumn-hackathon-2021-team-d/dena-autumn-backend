package entity

type Answer struct {
	ID         int    `json:"id" db:"id"`
	GroupID    string `json:"group_id" db:"group_id"`
	QuestionID int    `json:"question_id" db:"question_id"`
	Contents   string `json:"contents" db:"contents"`
	Username   string `json:"username" db:"username"`
	CreatedAt  string `json:"created_id" db:"created_at"`
}
