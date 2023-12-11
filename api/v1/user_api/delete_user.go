package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

func (UserApi) DeleteUser(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	id, _ := strconv.Atoi(idStr)
	var user models.User
	user.ID = uint(id)
	if err := user.DeleteByID(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "删除成功")
}
