package component_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils"
)

type DeleteComponentsRequest struct {
	IDs []int64 `json:"ids" validate:"required" label:"ids"`
}

func (ComponentApi) DeleteComponents(c *gin.Context) {
	var req DeleteComponentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if len(req.IDs) == 0 {
		response.FailWithMsg(c, response.FAIL_VALIDATE, "请选择")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	componentService := service.AppService.ComponentService
	if err := componentService.BatchDeleteComponentAndModify(req.IDs); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error())
		return
	}
	response.OkWithMsg(c, "删除成功")
}
