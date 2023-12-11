package jwt

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			global.Logrus.Error("未携带token")
			response.FailWithMsg(c, response.INVALID_PARAMS, "未携带token")
			c.Abort()
			return
		}
		// 解析token
		CustomClaims, err := jwts.ParseToken(token)
		if err != nil {
			response.FailWithMsg(c, response.FAIL_AUTH, "token验证失败")
			c.Abort()
			return
		}
		// 检查用户是否已经注销, redis中是否存在
		key := fmt.Sprintf("jwt_token:%d:%s", CustomClaims.Payload.UserID, token)
		exists, err := global.RedisClient.Exists(context.Background(), key).Result()
		if err != nil {
			global.Logrus.Error("token获取失败")
			response.FailWithMsg(c, response.FAIL_AUTH, "token获取失败")
			c.Abort()
		}
		if exists != 1 {
			global.Logrus.Error("登陆过期, 请重新登陆")
			response.FailWithMsg(c, response.FAIL_AUTH, "登陆过期, 请重新登陆")
			c.Abort()
		}
		c.Set("user", &CustomClaims.Payload)
		c.Next()
	}
}
