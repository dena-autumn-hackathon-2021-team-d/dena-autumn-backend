package controller

import (
	"encoding/json"
	"io"
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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to parse request body")
		a.logger.Errorf("failed to parse request body: :%v\n", err)
		return
	}

	var answer *entity.Answer
	if err := json.Unmarshal(body, answer); err != nil {
		c.String(http.StatusInternalServerError, "failed to unmarshal request body")
		a.logger.Errorf("failed to unmarshal request body: :%v\n", err)
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
