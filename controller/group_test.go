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

func TestGroup(t *testing.T) {
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

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

    reqBody := `{"name":"groupname"}`

    context.Request = httptest.NewRequest("GET", "/", bytes.NewBufferString(reqBody))

    groupCtrl.Create(context)

    var got entity.Group
	if err = json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err, string(w.Body.Bytes()))
	}

    want := entity.Group{
        Name:"groupname",
    }
    
    opts := cmpopts.IgnoreFields(got, "CreatedAt", "ID")
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}

