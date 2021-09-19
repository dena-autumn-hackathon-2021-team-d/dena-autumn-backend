package controller_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/controller"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
)

type fakeAnswerRepository struct {
	postFunc          func(answer *entity.Answer) error
	findByGroupIDFunc func(groupID string) ([]*entity.Answer, error)
	findUniqueFunc    func(groupID, questionID, answerID string) (*entity.Answer, error)
	findByQuestion    func(groupID, questionID string) ([]*entity.Answer, error)
}

func (a *fakeAnswerRepository) Post(answer *entity.Answer) error {
	return a.postFunc(answer)
}

func (a *fakeAnswerRepository) FindByGroupID(groupID string) ([]*entity.Answer, error) {
	return a.findByGroupIDFunc(groupID)
}

func (a *fakeAnswerRepository) FindUnique(groupID, questionID, answerID string) (*entity.Answer, error) {
	return a.findUniqueFunc(groupID, questionID, answerID)
}

func (a *fakeAnswerRepository) FindByQuestion(groupID, questionID string) ([]*entity.Answer, error) {
	return a.findByQuestion(groupID, questionID)
}

func TestAnswer_GetByGroupID(t *testing.T) {
	wantAnswers := []*entity.Answer{
		{
			ID:         "ID1",
			GroupID:    "GROUP_ID",
			QuestionID: "QUESTION_ID1",
			Contents:   "TEST_CONTENTS1",
			Username:   "TEST_USERNAME1",
			CreatedAt:  time.Date(2021, time.September, 18, 1, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
		{
			ID:         "ID2",
			GroupID:    "GROUP_ID",
			QuestionID: "QUESTION_ID2",
			Contents:   "TEST_CONTENTS2",
			Username:   "TEST_USERNAME2",
			CreatedAt:  time.Date(2021, time.September, 18, 2, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
	}

	answerRepo := &fakeAnswerRepository{
		findByGroupIDFunc: func(groupID string) ([]*entity.Answer, error) {
			return wantAnswers, nil
		},
	}
	answerUC := usecase.NewAnswerUseCase(answerRepo)
	answerCtrl := controller.NewAnswerController(log.New(), answerUC)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "group_id", Value: "GROUP_ID"})

	answerCtrl.GetByGroupID(c)

	var gotAnswers []*entity.Answer
	if err := json.Unmarshal(w.Body.Bytes(), &gotAnswers); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if diff := cmp.Diff(wantAnswers, gotAnswers); diff != "" {
		t.Fatal(diff)
	}
}
