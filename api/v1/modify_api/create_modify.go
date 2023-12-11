package modify_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type CreateModifyRequest struct {
	Title       string `json:"title" validate:"required" label:"标题"`
	Num         string `json:"num" validate:"required" label:"改装号"`
	Version     string `json:"version" validate:"required" label:"版本号"`
	ComponentID uint   `json:"component_id" validate:"required" label:"附件ID"`
}

func (ModifyApi) CreateModify(c *gin.Context) {
	var req CreateModifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var component models.Component
	if err := global.DB.Where("id = ?", req.ComponentID).Take(&component).Error; err != nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "附件不存在")
		return
	}

	modify := models.Modify{
		Title:       req.Title,
		Num:         req.Num,
		Version:     req.Version,
		ComponentID: req.ComponentID,
	}
	if err := modify.AddOneRecord(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "添加成功")
}
