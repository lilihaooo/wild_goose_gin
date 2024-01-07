package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

func (UserApi) GetUserRoleIDs(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	id, _ := strconv.Atoi(idStr)
	var model models.User
	model.ID = uint(id)
	user, err := model.TakeOneUserWithRoleIDs()
	if err != nil {
		if err != nil {
			global.Logrus.Error(err)
			response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
			return
		}
	}
	roleIDs := []uint{}
	if len(user.Roles) > 0 {
		for _, role := range user.Roles {
			roleIDs = append(roleIDs, role.ID)
		}
	}
	response.OkWithData(c, roleIDs)
}
