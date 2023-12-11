package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
)

type BindEmailValidateRequest struct {
	Email string `json:"email" validate:"required,email" label:"邮箱"`
	Code  string `json:"code" validate:"required" label:"验证码"`
}

func (UserApi) BindEmailValidate(c *gin.Context) {
	var cr BindEmailValidateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&cr); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	user := c.MustGet("user")
	payload := user.(*jwts.Payload)
	// 验证email
	hashName := "code:" + strconv.Itoa(int(payload.UserID))
	hashFields, err := global.RedisClient.HGetAll(c, hashName).Result()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "redis操作失败")
		return
	}
	if hashFields["code"] == "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE_CODE, "code已过期")
		return
	}

	if hashFields["code"] != cr.Code {
		response.FailWithMsg(c, response.FAIL_VALIDATE_CODE, "验证码不正确")
		return
	}
	if hashFields["email"] != cr.Email {
		response.FailWithMsg(c, response.FAIL_VALIDATE_CODE, "email不一致!!!")
		return
	}

	var userModel models.User
	userModel.ID = payload.UserID
	newData := map[string]interface{}{
		"email": cr.Email,
	}
	if row := global.DB.Model(&userModel).Updates(newData).RowsAffected; row == 0 {
		response.FailWithMsg(c, response.FAIL_OPER, "数据更新失败")
		return
	}
	response.OkWithMsg(c, "")
}
