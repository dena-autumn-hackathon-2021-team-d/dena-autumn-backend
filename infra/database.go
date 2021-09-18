  
package infra

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/config"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type DbMap *gorp.DbMap

// NewDB はSQLiteサーバに接続して、DbMapを生成します
func NewDB() (DbMap, error) {
	db, err := sql.Open("sqlite3", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite: %w", err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	logger := log.New()

	for {
		err := db.Ping()
		if err == nil {
			break
		}
		logger.Infof("%s\n", err.Error())
		time.Sleep(time.Second * 2)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	logger.Info("DB Ready!")
	return dbMap, nil
}
