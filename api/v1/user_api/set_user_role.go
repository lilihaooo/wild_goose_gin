package user_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type SetUserRoleReq struct {
	UserID  uint   `json:"user_id" validate:"required" label:"用户ID"`
	RoleIDs []uint `json:"role_ids" validate:"required" label:"角色IDs"`
}

func (UserApi) SetUserRole(c *gin.Context) {
	var req SetUserRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	// 验证用户是否存在
	var userModel models.User
	userModel.ID = req.UserID
	user, err := userModel.TakeOneUser()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}

	//验证角色是否存在  (过滤)
	var roleModel models.Role
	roles, err := roleModel.GetRolesByIDs(req.RoleIDs)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	if len(roles) == 0 {
		response.FailWithMsg(c, response.INVALID_PARAMS, "证书id不能为空")
		return
	}
	user.Roles = roles
	err = user.ReplaceRoles()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error())
		return
	}
	response.OkWithMsg(c, "添加成功")
}
