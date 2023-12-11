package modify_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

func (ModifyApi) DeleteModify(c *gin.Context) {
	IDStr := c.Query("id")
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	ID, _ := strconv.Atoi(IDStr)
	var model models.Modify
	model.ID = uint(ID)
	if err := model.DeleteOneRecord(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "删除成功")
}
