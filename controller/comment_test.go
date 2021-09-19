package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/controller"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type fakeCommentRepository struct {
	postFunc         func(comment *entity.Comment) error
	findUniqueFunc   func(groupID, questionID, answerID, commentID string) (*entity.Comment, error)
	findByAnswerFunc func(groupID, questionID, answerID string) ([]*entity.Comment, error)
}

func (f *fakeCommentRepository) Post(comment *entity.Comment) error {
	return f.postFunc(comment)
}

func (f *fakeCommentRepository) FindUnique(groupID, questionID, answerID, commentID string) (*entity.Comment, error) {
	return f.findUniqueFunc(groupID, questionID, answerID, commentID)
}

func (f *fakeCommentRepository) FindByAnswer(groupID, questionID, answerID string) ([]*entity.Comment, error) {
	return f.findByAnswerFunc(groupID, questionID, answerID)
}

func TestComment_Post(t *testing.T) {

	logger := log.New()

	dbMap, err := infra.NewDB()
	if err != nil {
		logger.Errorf("failed NewDB: %s", err.Error())
		os.Exit(1)
	}
	defer func() {
		err := dbMap.Db.Close()
		if err != nil {
			logger.Errorf("failed to close DB: %s", err.Error())
		}
	}()

	groupRepo := infra.NewGroupRepository(dbMap)
	groupUC := usecase.NewGroupUseCase(groupRepo)
	groupCtrl := controller.NewGroupController(logger, groupUC)

	answerRepo := infra.NewAnswerRepository(dbMap)
	answerUC := usecase.NewAnswerUseCase(answerRepo)
	answerCtrl := controller.NewAnswerController(logger, answerUC)

	questionRepo := infra.NewQuestionRepository(dbMap)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	questionCtrl := controller.NewQuestionController(logger, questionUC)

	commentRepo := infra.NewCommentRepository(dbMap)
	commentUC := usecase.NewCommentUseCase(commentRepo)
	commentCtrl := controller.NewCommentController(logger, commentUC)

	req := `{"name":"name"}`
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(req))
	groupCtrl.Create(context)

	var group entity.Group
	if err = json.Unmarshal(w.Body.Bytes(), &group); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	req = `{
		"contents":"Question?",
		"username":"user",
		"group_id":"` + group.ID + `",
	}`

	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(req))
	questionCtrl.Post(context)
	var question entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &question); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	req = `{
		"contents":"answer!",
		"username":"user",
		"group_id":` + group.ID + `",
		"question_id": "%s",
	}`
	req = fmt.Sprintf(req, question.ID)

	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(req))
	answerCtrl.Post(context)
	var answer entity.Answer
	if err = json.Unmarshal(w.Body.Bytes(), &answer); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	req = `{
		"contents":"comment!",
		"username":"user",
		"group_id":` + group.ID + `",
		"question_id": "%s",
		"answer_id": "%s",
	}`
	req = fmt.Sprintf(req, question.ID, answer.ID)

	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(req))
	commentCtrl.Post(context)

	var comment entity.Comment
	if err = json.Unmarshal(w.Body.Bytes(), &comment); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	want := entity.Comment{
		Contents:   "comment!",
		Username:   "user",
		GroupID:    group.ID,
		QuestionID: question.ID,
		AnswerID:   answer.ID,
	}

	opts := cmpopts.IgnoreFields(comment, "CreatedAt", "ID")
	if diff := cmp.Diff(want, comment, opts); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}

func TestComment_GetUnique(t *testing.T) {
	wantComment := entity.Comment{
		ID:         "1",
		GroupID:    "GroupID",
		QuestionID: "1",
		AnswerID:   "1",
		Contents:   "contents",
		Username:   "user",
		CreatedAt:  time.Date(2021, time.September, 18, 1, 0, 0, 0, time.UTC).Format(time.RFC3339),
	}

	commentRepo := &fakeCommentRepository{
		findUniqueFunc: func(groupID, questionID, answerID, commentID string) (*entity.Comment, error) {
			return &wantComment, nil
		},
	}

	commentUC := usecase.NewCommentUseCase(commentRepo)
	commentCtrl := controller.NewCommentController(log.New(), commentUC)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "group_id", Value: "GroupID"}, gin.Param{Key: "question_id", Value: "1"}, gin.Param{Key: "answer_id", Value: "1"}, gin.Param{Key: "comment_id", Value: "1"})
	commentCtrl.GetUnique(c)
	var comment entity.Comment
	if err := json.Unmarshal(w.Body.Bytes(), &comment); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if diff := cmp.Diff(wantComment, comment); diff != "" {
		t.Fatal(diff)
	}
}

func TestComment_GetByAnswer(t *testing.T) {
	wantComment := []*entity.Comment{
		{
			ID:         "1",
			GroupID:    "GroupID",
			QuestionID: "1",
			AnswerID:   "1",
			Contents:   "contents",
			Username:   "user",
			CreatedAt:  time.Date(2021, time.September, 18, 1, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
		{
			ID:         "2",
			GroupID:    "GroupID",
			QuestionID: "1",
			AnswerID:   "1",
			Contents:   "contents",
			Username:   "user",
			CreatedAt:  time.Date(2021, time.September, 18, 1, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
	}

	commentRepo := &fakeCommentRepository{
		findByAnswerFunc: func(groupID string, questionID string, answerID string) ([]*entity.Comment, error) {
			return wantComment, nil
		},
	}

	commentUC := usecase.NewCommentUseCase(commentRepo)
	commentCtrl := controller.NewCommentController(log.New(), commentUC)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "group_id", Value: "GroupID"}, gin.Param{Key: "question_id", Value: "1"}, gin.Param{Key: "answer_id", Value: "1"})
	commentCtrl.GetByAnswer(c)

	var comment []*entity.Comment
	if err := json.Unmarshal(w.Body.Bytes(), &comment); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if diff := cmp.Diff(wantComment, comment); diff != "" {
		t.Fatal(diff)
	}
}
