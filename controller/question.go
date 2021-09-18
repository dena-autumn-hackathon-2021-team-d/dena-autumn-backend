package controller

import (
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

type QuestionController struct {
	uc     *usecase.QuestionUsecase
	logger *log.Logger
}

func NewQuestionController(logger *log.Logger, uc *usecase.QuestionUsecase) *QuestionController {
	return &QuestionController{uc: uc, logger: logger}
}

func (ctrl *QuestionController) Post(c *gin.Context) {
	var question *entity.Question
	if err := c.ShouldBindJSON(question); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		ctrl.logger.Errorf("failed to bind request body: :%v\n", err)
		return
	}

	resQuestion, err := ctrl.uc.Post(question)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post answer")
		ctrl.logger.Errorf("failed to post qestion: :%v\n", err)
		return
	}

	c.JSON(http.StatusOK, resQuestion)
}
