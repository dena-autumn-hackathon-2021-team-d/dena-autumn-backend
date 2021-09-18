package main

import "github.com/gin-gonic/gin"

func main() {
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