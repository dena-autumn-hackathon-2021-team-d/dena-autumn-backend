package controller

import (
	"net/http"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/log"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	uc *usecase.QuestionUsecase
}

func NewQuestionController (uc *usecase.QuestionUsecase) *QuestionController{
	return &QuestionController{uc : uc}
}

func(ctrl *QuestionController) Post (c gin.Context) {
	logger := log.New()
	var question *entity.Question
	if err := c.ShouldBindJSON(question); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		logger.Errorf("failed to bind request body: :%v\n", err)
		return
	}

	resQuestion, err := ctrl.uc.Post(question)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post answer")
		logger.Errorf("failed to post qestion: :%v\n", err)
		return
	}

	c.JSON(http.StatusOK, resQuestion)
}