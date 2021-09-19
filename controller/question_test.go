package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/controller"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestQuestion(t *testing.T) {
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

	questionRepo := infra.NewQuestionRepository(dbMap)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	questionCtrl := controller.NewQuestionController(logger, questionUC)

	answerRepo := infra.NewAnswerRepository(dbMap)
	answerUC := usecase.NewAnswerUseCase(answerRepo)
	answerCtrl := controller.NewAnswerController(logger, answerUC)

	// Groupの作成
	reqBody := `{"name":"groupname"}`
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(reqBody))
	groupCtrl.Create(context)

	var group entity.Group
	if err = json.Unmarshal(w.Body.Bytes(), &group); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	// Questionの作成
	reqBody = `{
		"contents":"Question1",
		"username":"user",
		"group_id":"` + group.ID + `"
	}`

	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(reqBody))
	questionCtrl.Post(context)
	var question entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &question); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}
	want := entity.Question{
		Contents:   "Question1",
		Username:   "user",
		GroupID:    group.ID,
		NumAnswers: 0,
	}
	opts := cmpopts.IgnoreFields(question, "CreatedAt", "ID")
	if diff := cmp.Diff(want, question, opts); diff != "" {
		t.Errorf("Post (-want +got) =\n%s\n", diff)
	}

	// num_answersの確認のためにanswersを作る
	reqBodys := []string{
		`{
			"contents":"ANSWERS1",
			"username":"user",
			"group_id":"` + group.ID + `",
			"question_id":"` + question.ID + `"
		}`,
		`{
			"contents":"ANSWERS2",
			"username":"user",
			"group_id":"` + group.ID + `",
			"question_id":"` + question.ID + `"
		}`,
	}
	for _, r := range reqBodys {
		w = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(w)
		context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(r))
		answerCtrl.Post(context)
	}

	// FindByQuestionが正しく取得できる
	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", nil)
	context.Params = append(context.Params,
		gin.Param{Key: "group_id", Value: group.ID},
		gin.Param{Key: "question_id", Value: question.ID},
	)
	questionCtrl.FindByQuestion(context)
	var got entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}
	question.NumAnswers = 2
	if diff := cmp.Diff(question, got); diff != "" {
		t.Errorf("FindByQuestion (-want +got) =\n%s\n", diff)
	}

	// Questionの２つ目を追加する
	reqBody = `{
		"contents":"Question2",
		"username":"user",
		"group_id":"` + group.ID + `"
	}`

	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(reqBody))
	questionCtrl.Post(context)
	var question2 entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &question2); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	// GetAllが正しく取得できる
	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", nil)
	context.Params = append(context.Params,
		gin.Param{Key: "group_id", Value: group.ID},
	)
	questionCtrl.GetAll(context)
	var gotQuestions []entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &gotQuestions); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

	wantQuestions := []entity.Question{
		question,
		question2,
	}
	if diff := cmp.Diff(wantQuestions, gotQuestions); diff != "" {
		t.Errorf("GetAll (-want +got) =\n%s\n", diff)
	}

}
