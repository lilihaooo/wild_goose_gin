package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
	"wild_goose_gin/utils/u_redis"
)

type updateRoleRequest struct {
	ID       uint   `json:"id" validate:"required"`
	RoleID   uint   `validate:"required"`
	NickName string `json:"nick_name" validate:"omitempty,max=36"` // 当管理员认为昵称不合法时会修改其昵称
}

func (UserApi) UserUpdateRole(c *gin.Context) {
	// 获取用户角色
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	if payload.RoleID != common_type.Admin {
		response.FailWithMsg(c, response.ERROR_AUTH_CHECK_FAIL, "")
		return
	}
	var req updateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	// 1号用户永远都是超级管理员
	if req.ID == common_type.Admin {
		response.FailWithMsg(c, response.FAIL_OPER, "超级管理员, 不允许修改")
		return
	}

	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	// 查询是否存在
	var model models.User
	err := global.DB.Take(&model, req.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailWithMsg(c, response.ERROR_NOT_EXIST_USER, "")
			return
		}
		global.Logrus.Error(err)
		return
	}

	data := map[string]any{
		"role_id":   req.RoleID,
		"nick_name": req.NickName,
	}
	if err = global.DB.Model(&model).Updates(data).Error; err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "修改角色失败")
		return
	}

	// 删除该用户所有token
	matchKey := fmt.Sprintf("jwt_token:%d:*", model.ID)
	u_redis.DeleteAllKeys(global.RedisClient, matchKey)
	response.OkWithMsg(c, "")
}
