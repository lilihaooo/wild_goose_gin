package role_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

func (RoleApi) GetAllRoleList(c *gin.Context) {
	var model models.Role
	roles, err := model.GetAllRoleList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithData(c, roles)
}
