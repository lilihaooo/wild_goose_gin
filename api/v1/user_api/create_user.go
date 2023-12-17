package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/utils"

	"wild_goose_gin/pkg/response"
)

type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required" label:"用户名"`
	RoleID   uint   `json:"role_id" validate:"required" label:"角色ID"`
}

func (UserApi) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var model models.User
	if err := global.DB.Where("user_name = ?", req.UserName).Take(&model).Error; err == nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "用户名已存在")
		return
	}

	// 对密码进行加密处理
	hashPassword, err := utils.HashPassword("111111") // 初始密码
	if err != nil {
		fmt.Println("密码加密失败!!!")
		return
	}
	model = models.User{
		UserName: req.UserName,
		RoleID:   req.RoleID,
		Password: hashPassword,
	}
	if err := model.AddOneRecord(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "添加成功")

}
