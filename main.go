package main

import (
	"os"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/controller"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
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

	groupRepo := infra.NewGroupRepository(dbMap)
	groupUC := usecase.NewGroupUseCase(groupRepo)
	groupCtrl := controller.NewGroupController(logger, groupUC)

	answerRepo := infra.NewAnswerRepository(dbMap)
	answerUC := usecase.NewAnswerUseCase(answerRepo)
	answerCtrl := controller.NewAnswerController(logger, answerUC)

	r := gin.Default()

	api := r.Group("/api")

	//グループの作成
	api.POST("/group", groupCtrl.Create)
	//質問の作成
	api.POST("/question", func(c *gin.Context) {})
	//該当する質問を取得する
	api.GET("/group/:group_id/question/:question_id", func(c *gin.Context) {})
	//グループの質問一覧
	api.GET("/group/:group_id/questions", func(c *gin.Context) {})

	//解答のポスト
	api.POST("/answer", answerCtrl.Post)
	//グループ全体の解答一覧を取得する
	api.GET("/group/:group_id/answers", answerCtrl.GetByGroupID)
	//該当の答えを取得する
	api.GET("/group/:group_id/question/:question_id/answer/:answer_id", func(c *gin.Context) {})

	//もしかしたら削る
	//コメントの投稿
	api.POST("/comment")
	api.GET("/group/:group_id/question/:question_id/answer/:answer_id/comment/:comment_id")

	r.Run(":8000")
}
