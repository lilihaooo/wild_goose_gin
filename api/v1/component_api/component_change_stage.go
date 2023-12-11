package component_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

func (ComponentApi) ComponentChangeState(c *gin.Context) {
	IDStr := c.Query("id")
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	ID, _ := strconv.Atoi(IDStr)

	var component models.Component
	component.Model.ID = uint(ID)
	err := component.ChangeState()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "")
}
