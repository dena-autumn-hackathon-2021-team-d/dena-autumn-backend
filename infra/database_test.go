package infra_test

import (
	"testing"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
)

func TestNewDB(t *testing.T) {
	dbMap, err := infra.NewDB()
	if err != nil {
		t.Errorf("failed NewDB: %s", err.Error())
	}

	// きちんとCloseすること
	defer func() {
		err := dbMap.Db.Close()
		if err != nil {
			t.Errorf("failed to close DB: %s", err.Error())
		}
	}()
}
