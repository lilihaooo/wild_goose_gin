package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
	"wild_goose_gin/global"
)

type Payload struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	RoleID   uint   `json:"role_id"`
}

type CustomClaims struct {
	Payload
	jwt.StandardClaims
}

// GetJwtExpiresDuration 获得jwt过期时间就间隔
func GetJwtExpiresDuration() time.Duration {
	return time.Minute * time.Duration(global.Config.Jwt.Expires)
}
