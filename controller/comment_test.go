package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	questionRepo := infra.NewQuestionRepository(dbMap)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	questionCtrl := controller.NewQuestionController(logger, questionUC)

	answerRepo := infra.NewAnswerRepository(dbMap)
	answerUC := usecase.NewAnswerUseCase(answerRepo)
	answerCtrl := controller.NewAnswerController(logger, answerUC)

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
		"group_id":"`+group.ID+`"
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
		"group_id":`+group.ID+`",
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
		"group_id":`+group.ID+`",
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
		Contents: "comment!",
		Username: "user",
		GroupID: group.ID,
		QuestionID: question.ID,
		AnswerID: answer.ID,
	}

	opts := cmpopts.IgnoreFields(comment, "CreatedAt", "ID")
	if diff := cmp.Diff(want, comment, opts); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}
