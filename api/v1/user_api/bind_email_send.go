package user_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
	"wild_goose_gin/utils/u_email"
	"wild_goose_gin/utils/u_random"
)

type BindEmailSendRequest struct {
	Email string `json:"email" validate:"required,email" label:"邮箱"`
}

func (UserApi) BindEmailSend(c *gin.Context) {
	var cr BindEmailSendRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&cr); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var userModel models.User
	if row := global.DB.Where("email = ?", cr.Email).Find(&userModel).RowsAffected; row > 0 {
		response.FailWithMsg(c, response.INVALID_PARAMS, "该email已被绑定")
		return
	}

	//发送验证码
	code := u_random.GenRandomCode(4)
	if err := u_email.SendEmail("验证码", code, cr.Email); err != nil {
		response.FailWithMsg(c, response.FAIL_SEND_EMAIl, "")
		return
	}
	user := c.MustGet("user")
	payload := user.(*jwts.Payload)

	// 将id为key, code和email为字段以hash类型保存到redis中,同时过期时间为60s 利用lua脚本保证添加与设置过期时间的原子性
	hashName := "code:" + strconv.FormatUint(uint64(payload.UserID), 10) // 将 uint 转换为 string
	// lua脚本的参数格式{过期时间(s), key1, value1, key2, value2, ...}
	args := []interface{}{6000, "code", code, "email", cr.Email} // 参数列表
	err := service.AppService.SetHashWithExpireTime(hashName, args)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "redis操作失败"+err.Error())
		return
	}
	response.OkWithMsg(c, "")
}
