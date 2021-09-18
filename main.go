package main

import (
	"os"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/gin-gonic/gin"
)

func main() {
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
	
	r := gin.Default()

	// 質問をランダムで取得する
	r.GET("/api/question", func(c *gin.Context){})

	//タイムラインの一覧
	r.GET("/api/answers", func(c *gin.Context) {})

	//回答をポスト
	r.POST("/api/answer", func(c *gin.Context) {})

	r.GET("/api/ansewer/:answer_id", func(c *gin.Context) {})
	r.Run(":8000")
}
