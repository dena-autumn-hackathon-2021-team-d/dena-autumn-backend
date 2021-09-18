package controller

import (
	"net/http"

	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"
	"github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type GroupController struct {
	logger   *log.Logger
	uc *usecase.GroupUseCase
}

func NewGroupController(logger *log.Logger, uc *usecase.GroupUseCase) *GroupController {
	return &GroupController{
        logger: logger,
        uc: uc,
    }
}

func (ctrl *GroupController) Create(c *gin.Context) {
    var group *entity.Group 
    err := c.ShouldBindJSON(&group)
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
