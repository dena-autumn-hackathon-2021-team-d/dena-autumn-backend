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

	//グループの作成
	r.POST("/group", func(c *gin.Context) {})
	//質問の作成
	r.POST("/question", func(c *gin.Context) {})
	//該当する質問を取得する
	r.GET("/group/:group_id/question/:id", func(c *gin.Context) {})
	//グループの質問一覧
	r.GET("/group/:group_id/questions", func(c *gin.Context) {})

	//解答のポスト
	r.POST("/answer", func(c *gin.Context) {})
	//該当の答えを取得する
	r.GET("/group/:group_id/question/:question_id/answer/:answer_id", func(c *gin.Context) {})

	//もしかしたら削る
	//コメントの投稿
	r.POST("/commnet")
	r.GET("/group/:group_id/question/:question_id/answer/:answer_id/comment/:comment_id")
	r.Run(":8000")
}
