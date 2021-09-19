package controller

import (
	"net/http"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
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
	if err := c.ShouldBindJSON(&question); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		ctrl.logger.Errorf("failed to bind request body: :%v\n", err)
		return
	}

	resQuestion, err := ctrl.uc.Post(question)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post answer")
		ctrl.logger.Errorf("failed to post question: :%v\n", err)
		return
	}

	c.JSON(http.StatusOK, resQuestion)
}

func (ctrl *QuestionController) GetRandomly(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter group_id")
		ctrl.logger.Errorf("invalid path parameter group_id: group_id=%s", groupID)
		return
	}

	question, err := ctrl.uc.GetRandomly(groupID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get question Randomly")
		ctrl.logger.Errorf("failed to get question Randomly: %v", err)
		return
	}

	c.JSON(http.StatusOK, question)
}

func (ctrl *QuestionController) FindByQuestion(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter group_id")
		ctrl.logger.Errorf("invalid path parameter group_id: group_id=%s", groupID)
		return
	}

	questionID := c.Param("question_id")
	if questionID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter question_id")
		ctrl.logger.Errorf("invalid path parameter question_id: question_id=%s", questionID)
		return
	}

	resQuestion, err := ctrl.uc.FindByQuestion(groupID, questionID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get question")
		ctrl.logger.Errorf("failed to get by question: %v", err)
		return
	}

	c.JSON(http.StatusOK, resQuestion)
}

func (ctrl *QuestionController) GetAll(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter group_id")
		ctrl.logger.Errorf("invalid path parameter group_id: group_id=%s", groupID)
		return
	}

	resQuestions, err := ctrl.uc.GetAll(groupID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get all questions")
		ctrl.logger.Errorf("failed to get by all questions: %v", err)
		return
	}

	c.JSON(http.StatusOK, resQuestions)
}
