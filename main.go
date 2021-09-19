package main

import (
	"os"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/config"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/controller"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/infra"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-contrib/cors"
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

	questionRepo := infra.NewQuestionRepository(dbMap)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	questionCtrl := controller.NewQuestionController(logger, questionUC)

	commentRepo := infra.NewCommentRepository(dbMap)
	commentUC := usecase.NewCommentUseCase(commentRepo)
	commentCtrl := controller.NewCommentController(logger, commentUC)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowCredentials: true,
	}))

	api := r.Group("/api")

	//グループの作成
	api.POST("/group", groupCtrl.Create)
	//該当するグループの情報を取得する
	api.GET("/group/:group_id", groupCtrl.GetByID)

	//質問の作成
	api.POST("/question", questionCtrl.Post)
	//ランダムに質問を取得する
	api.GET("/group/:group_id/question", questionCtrl.GetRandomly)
	//該当する質問を取得する
	api.GET("/group/:group_id/question/:question_id", questionCtrl.FindByQuestion)
	//グループの質問一覧
	api.GET("/group/:group_id/questions", questionCtrl.GetAll)

	//解答のポスト
	api.POST("/answer", answerCtrl.Post)
	//グループ全体の解答一覧を取得する
	api.GET("/group/:group_id/answers", answerCtrl.GetByGroupID)
	//該当の答えを取得する
	api.GET("/group/:group_id/question/:question_id/answer/:answer_id", answerCtrl.GetUnique)
	//質問に紐付いた解答一覧を取得する
	api.GET("/group/:group_id/question/:question_id/answers", answerCtrl.GetByQuestion)

	//コメントの投稿
	api.POST("/comment", commentCtrl.Post)
	//該当のコメントを取得する
	api.GET("/group/:group_id/question/:question_id/answer/:answer_id/comment/:comment_id", commentCtrl.GetUnique)
	//コメント一覧を取得する
	api.GET("/group/:group_id/question/:question_id/answer/:answer_id/comments", commentCtrl.GetByAnswer)

	r.Run(config.Port())
}
