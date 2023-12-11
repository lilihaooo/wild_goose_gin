package user_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
)

type updatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	// 获取用户角色
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	// 鉴权
	if payload.RoleID != common_type.Admin {
		response.FailWithMsg(c, response.ERROR_AUTH_CHECK_FAIL, "")
		return
	}
	var cr updatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&cr); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	if cr.Password == cr.OldPassword {
		response.FailWithMsg(c, response.INVALID_PARAMS, "两次输入的密码一致")
		return
	}

	// 查询user是否存在
	var model models.User
	err := global.DB.Take(&model, payload.UserID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailWithMsg(c, response.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}
	// 对旧密码进行验证
	if ok := utils.CheckPasswordHash(cr.OldPassword, model.Password); !ok {
		response.FailWithMsg(c, response.FAIL_OPER, "密码错误!!!")
		return
	}
	// 对新密码进行加密处理
	hashPassword, err := utils.HashPassword(cr.Password)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "密码加密失败!!!")
		return
	}
	if err = global.DB.Model(&model).Update("password", hashPassword).Error; err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "修改密码失败")
		return
	}
	response.OkWithMsg(c, "添加成功")
}
