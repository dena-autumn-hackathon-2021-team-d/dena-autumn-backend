package usecase

import (
	"time"
	"fmt"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
)

type QuestionUsecase struct {
	questionRepo repository.Question
}

func NewQuestionUseCase(question repository.Question) *QuestionUsecase {
	return &QuestionUsecase{questionRepo: question}
}

func (q *QuestionUsecase) Post(question *entity.Question) (*entity.Question, error) {
	question.CreatedAt = time.Now().Format(time.RFC3339)
	if err := q.questionRepo.Post(question); err != nil {
		return nil, fmt.Errorf("failed to post answer: %w", err)
	}
	return question, nil
}
