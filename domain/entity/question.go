package entity

type Question struct {
	ID        string `json:"id" db:"id"`
	Contents  string `json:"contents" db:"contents"`
	GroupID   string `json:"group_id" db:"group_id"`
	Username  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
