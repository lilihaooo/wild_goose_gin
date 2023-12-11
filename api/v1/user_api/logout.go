package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils/jwts"
)

func (UserApi) Logout(c *gin.Context) {
	token := c.GetHeader("token")
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	key := fmt.Sprintf("jwt_token:%d:%s", payload.UserID, token)
	_, err := global.RedisClient.Del(c, key).Result()
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.FAIL_OPER, "注销失败")
		return
	}
	response.OkWithMsg(c, "")
}
