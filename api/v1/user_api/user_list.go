package user_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	serializer "wild_goose_gin/serialize"

	"wild_goose_gin/pkg/response"
)

func (UserApi) UserList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	var model models.User
	users, count, err := model.GetUserList(req)
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
		return
	}
	res := serializer.BuildUsers(users)
	response.OkWithList(c, res, count)
}
