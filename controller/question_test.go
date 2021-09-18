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

	// answerRepo := infra.NewAnswerRepository(dbMap)
	// answerUC := usecase.NewAnswerUseCase(answerRepo)
	// answerCtrl := controller.NewAnswerController(logger, answerUC)

	questionRepo := infra.NewQuestionRepository(dbMap)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	questionCtrl := controller.NewQuestionController(logger, questionUC)

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
    "contents":"Question?",
    "username":"user",
    "group_id":"`+group.ID+`"
}`

	fmt.Println(reqBody)
	
	w = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(w)
    context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(reqBody))
    questionCtrl.Post(context)

    var got entity.Question
	if err = json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

    want := entity.Question{
        Contents:"Question?",
		Username: "user",
		GroupID: group.ID,
    }
    
    opts := cmpopts.IgnoreFields(got, "CreatedAt", "ID")
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}
