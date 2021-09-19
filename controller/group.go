package controller

import (
	"net/http"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type GroupController struct {
	logger *log.Logger
	uc     *usecase.GroupUseCase
}

func NewGroupController(logger *log.Logger, uc *usecase.GroupUseCase) *GroupController {
	return &GroupController{
		logger: logger,
		uc:     uc,
	}
}

func (ctrl *GroupController) Create(c *gin.Context) {
	group := &entity.Group{}
	err := c.ShouldBindJSON(group)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		ctrl.logger.Errorf("failed to bind request body: :%v\n", err)
		return
	}

	if err := ctrl.uc.Create(group); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		ctrl.logger.Errorf("failed to create group: :%v\n", err)
		return
	}

	c.JSON(http.StatusCreated, group)
	return
}

func (ctrl *GroupController) GetByID(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.String(http.StatusBadRequest, "invalid path parameter group_id")
		ctrl.logger.Errorf("invalid path parameter group_id: group_id=%s", groupID)
		return
	}

	group, err := ctrl.uc.GetByID(groupID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		ctrl.logger.Errorf("failed to get group by id: %w", err)
		return
	}

	c.JSON(http.StatusOK, group)
}
