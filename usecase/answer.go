package usecase

import (
	"fmt"
	"time"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/repository"
)

type AnswerUseCase struct {
	answerRepo repository.Answer
}

func NewAnswerUseCase(answerRepo repository.Answer) *AnswerUseCase {
	return &AnswerUseCase{answerRepo: answerRepo}
}

func (a *AnswerUseCase) Post(answer *entity.Answer) (*entity.Answer, error) {
	answer.CreatedAt = time.Now().Format(time.RFC3339)

	if err := a.answerRepo.Post(answer); err != nil {
		return nil, fmt.Errorf("failed to post answer: %w", err)
	}

	return answer, nil
}