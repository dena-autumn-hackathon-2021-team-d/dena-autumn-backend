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
	return &AnswerController{logger: logger, answerUC: answerUC}
}

func (a *AnswerController) Post(c *gin.Context) {
	var answer *entity.Answer
	if err := c.ShouldBindJSON(answer); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		a.logger.Errorf("failed to bind request body: :%v", err)
		return
	}

	resAnswer, err := a.answerUC.Post(answer)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post answer")
		a.logger.Errorf("failed to post answer: :%v", err)
		return
	}

	c.JSON(http.StatusOK, resAnswer)
}

func (a *AnswerController) GetByGroupID(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter group_id")
		a.logger.Errorf("invalid path parameter group_id: group_id=%s", groupID)
		return
	}

	resAnswers, err := a.answerUC.GetByGroupID(groupID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get answers by group id")
		a.logger.Errorf("failed to get answers by group id: %v", err)
		return
	}

	c.JSON(http.StatusOK, resAnswers)
}
