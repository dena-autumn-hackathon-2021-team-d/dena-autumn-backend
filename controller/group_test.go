package controller_test

import (
	"bytes"
	"encoding/json"
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
		Name: "groupname",
	}

	opts := cmpopts.IgnoreFields(got, "CreatedAt", "ID")
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}

func TestGroup_GetByID(t *testing.T) {
	wantGroup := &entity.Group{
		ID:        "GROUP_ID",
		Name:      "GROUP_NAME",
		CreatedAt: time.Date(2021, time.September, 19, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
	}

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

	dbMap.AddTableWithName(entity.Group{}, "groups").SetKeys(false, "id")
	if err := dbMap.Insert(wantGroup); err != nil {
		t.Fatalf("failed to insert group: %v", err)
	}
	defer func() {
		if _, err := dbMap.Delete(wantGroup); err != nil {
			t.Fatalf("failed to delete group: %v", err)
		}
	}()

	groupRepo := infra.NewGroupRepository(dbMap)
	groupUC := usecase.NewGroupUseCase(groupRepo)
	groupCtrl := controller.NewGroupController(logger, groupUC)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "group_id", Value: "GROUP_ID"})

	groupCtrl.GetByID(c)

	gotGroup := &entity.Group{}
	if err := json.Unmarshal(w.Body.Bytes(), gotGroup); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if diff := cmp.Diff(wantGroup, gotGroup); diff != "" {
		t.Fatalf(diff)
	}
}
