package user_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"

	"wild_goose_gin/pkg/response"
)

type UpdateUserRequest struct {
	UserName string `json:"user_name" validate:"required" label:"用户名"`
	RoleID   uint   `json:"role_id" validate:"required" label:"角色ID"`
	GroupID  uint   `json:"group_id"`
}

func (UserApi) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var model models.User
	if err := global.DB.Where("user_name = ?", req.UserName).Take(&model).Error; err != nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "用户不存在")
		return
	}
	if model.RoleID == common_type.Admin && req.RoleID != common_type.Admin {
		response.FailWithMsg(c, response.FAIL_OPER, "管理员角色不能修改")
		return
	}
	model.RoleID = req.RoleID
	model.GroupID = &req.GroupID
	if err := model.UpdateOneRecord(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "修改成功")

}
