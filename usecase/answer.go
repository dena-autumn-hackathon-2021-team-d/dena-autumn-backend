package usecase

import (
	"fmt"
	"strconv"
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

func (a *AnswerUseCase) GetByGroupID(groupID string) ([]*entity.Answer, error) {
	return a.answerRepo.FindByGroupID(groupID)
}

func (a *AnswerUseCase) GetUnique(groupID, questionID, answerID string) (*entity.Answer, error) {
	iQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse question id: %w", err)
	}

	iAnswerID, err := strconv.Atoi(answerID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse answer id: %w", err)
	}

	return a.answerRepo.FindUnique(groupID, iQuestionID, iAnswerID)
}
