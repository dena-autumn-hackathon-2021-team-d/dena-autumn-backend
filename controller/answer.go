package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
)

type AnswerController struct {
	logger   *log.Logger
	answerUC *usecase.AnswerUseCase
}

func NewAnswerController(logger *log.Logger, answerUC *usecase.AnswerUseCase) *AnswerController {
	return &AnswerController{answerUC: answerUC}
}

func (a *AnswerController) Post(c *gin.Context) {
	var answer *entity.Answer
	if err := c.ShouldBindJSON(answer); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		a.logger.Errorf("failed to bind request body: :%v\n", err)
		return
	}

	resAnswer, err := a.answerUC.Post(answer)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post answer")
		a.logger.Errorf("failed to post answer: :%v\n", err)
		return
	}

	c.JSON(http.StatusOK, resAnswer)
}
