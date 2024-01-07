package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/serialize/user_serialize"
)

func (UserApi) GetUserAuthorizeList(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	id, _ := strconv.Atoi(idStr)
	var model models.UserManual
	model.UserID = uint(id)
	userManuals, err := model.GetRecordByUserID()
	if err != nil {
		if err != nil {
			global.Logrus.Error(err)
			response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
			return
		}
	}
	list := user_serialize.BuildUserAuthorizes(userManuals)
	response.OkWithData(c, list)
}
