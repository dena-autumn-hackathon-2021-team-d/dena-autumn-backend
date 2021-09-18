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
	var comment *entity.Comment
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
