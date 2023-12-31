package component_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type CreateComponentRequest struct {
	Name     string `json:"name" validate:"required" label:"件名"`
	PN       string `json:"pn" validate:"required" label:"件号"`
	ManualID uint   `json:"manual_id" validate:"required" label:"手册ID"`
	GroupID  uint   `json:"group_id" validate:"required" label:"小组"`
}

func (ComponentApi) CreateComponent(c *gin.Context) {
	var req CreateComponentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var component models.Component
	if err := global.DB.Where("name = ?", req.Name).Take(&component).Error; err == nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "件名已存在")
		return
	}
	if err := global.DB.Where("pn = ?", req.PN).Take(&component).Error; err == nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "件号已存在")
		return
	}
	var manual models.Manual
	if err := global.DB.Where("id = ?", req.ManualID).Take(&manual).Error; err != nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "手册不存在")
		return
	}
	component = models.Component{
		Name:     req.Name,
		PN:       req.PN,
		ManualID: req.ManualID,
		GroupID:  req.GroupID,
	}
	if err := component.AddOneRecord(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "添加成功")
}
