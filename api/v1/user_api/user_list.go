package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/serialize/user_serialize"
	"wild_goose_gin/service"

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
	res := user_serialize.BuildUsers(users)
	response.OkWithList(c, res, count)
}

func (UserApi) UserAllSelectList(c *gin.Context) {
	var model models.User
	users, err := model.GetAllUserSelectList()
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
		return
	}
	res := user_serialize.BuildUsersSelect(users)
	response.OkWithData(c, res)
}

func (UserApi) UserAuthorizeUserList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	var model models.User
	users, count, err := model.GetUserAuthorizeUserList(req)
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
		return
	}
	res := user_serialize.BuildUsersListForAuthorize(users)
	response.OkWithList(c, res, count)
}

// UserTaskOptionalList 获得某个任务可以选择的用户列表
func (UserApi) UserTaskOptionalList(c *gin.Context) {
	taskIDStr := c.Query("id")
	if taskIDStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	taskID, _ := strconv.Atoi(taskIDStr)

	groupIDStr := c.Query("group_id")
	if groupIDStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "group_id不能为空")
		return
	}
	groupID, _ := strconv.Atoi(groupIDStr)

	users, err := service.AppService.GetUserTaskOptionalList(uint(taskID), uint(groupID))
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error())
		return
	}

	res := user_serialize.BuildUsersSelect(users)
	response.OkWithData(c, res)
}
