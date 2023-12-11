package settings_api

import (
	"github.com/gin-gonic/gin"

	"wild_goose_gin/config"
	"wild_goose_gin/core"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/response"
)

func (SettingsApi) GetEmail(c *gin.Context) {
	email := global.Config.Email
	response.OkWithData(c, email)
}

func (SettingsApi) UpdateEmail(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.Email = cr
	if !core.SetYaml() {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "")
}
