package user_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
)

type LoginRequest struct {
	UserName string `json:"user_name"  validate:"max=36,required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

func (UserApi) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	// 查找用户是否为空
	var userModel models.User
	userModel.UserName = req.UserName
	user, err := userModel.GetUserInfoByUserName()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailWithMsg(c, response.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}
	// 验证密码是否正确
	ok := utils.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		response.FailWithMsg(c, response.ERROR_PASS_USER, "")
		return
	}
	payload := jwts.Payload{
		UserID:   user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		RoleID:   user.Role.ID,
	}
	token := jwts.GenToken(payload)
	key := fmt.Sprintf("jwt_token:%d:%s", user.ID, token)
	// 将jwt保存到redis中
	global.RedisClient.Set(context.Background(), key, "", jwts.GetJwtExpiresDuration())
	data := map[string]any{
		"token": token,
	}
	response.OkWithData(c, data)
}
