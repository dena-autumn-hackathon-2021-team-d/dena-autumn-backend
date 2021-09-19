package controller

import (
	"net/http"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type CommentController struct {
	logger    *log.Logger
	commentUC *usecase.CommentUseCase
}

func NewCommentController(logger *log.Logger, commentUC *usecase.CommentUseCase) *CommentController {
	return &CommentController{
		logger:    logger,
		commentUC: commentUC,
	}
}

func (ctrl *CommentController) Post(c *gin.Context) {
	comment := &entity.Comment{}
	if err := c.ShouldBindJSON(comment); err != nil {
		c.String(http.StatusInternalServerError, "failed to bind request body")
		ctrl.logger.Errorf("failed to bind request body: %v", err)
		return
	}

	resComment, err := ctrl.commentUC.Post(comment)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to post comment")
		ctrl.logger.Errorf("failed to post comment: %v", err)
		return
	}

	c.JSON(http.StatusOK, resComment)
}

func (ctrl *CommentController) GetUnique(c *gin.Context) {
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

	answerID := c.Param("answer_id")
	if answerID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter answer_id")
		ctrl.logger.Errorf("invalid path parameter answer_id: answer_id=%s", answerID)
		return
	}

	commentID := c.Param("comment_id")
	if commentID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter comment_id")
		ctrl.logger.Errorf("invalid path parameter comment_id: comment_id=%s", commentID)
		return
	}

	comment, err := ctrl.commentUC.GetUnique(groupID, questionID, answerID, commentID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get unique comment")
		ctrl.logger.Errorf("failed to get unique comment: %v", err)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (ctrl *CommentController) GetByAnswer(c *gin.Context) {
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

	answerID := c.Param("answer_id")
	if answerID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter answer_id")
		ctrl.logger.Errorf("invalid path parameter answer_id: answer_id=%s", answerID)
		return
	}

	comments, err := ctrl.commentUC.GetByAnswer(groupID, questionID, answerID)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get comments by answer")
		ctrl.logger.Errorf("failed to get comments by answer: %v", err)
		return
	}

	c.JSON(http.StatusOK, comments)
}
