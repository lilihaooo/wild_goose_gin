package settings_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/config"
	"wild_goose_gin/core"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/response"
)

func (SettingsApi) GetQiNiu(c *gin.Context) {
	qiniu := global.Config.QiNiu
	response.OkWithData(c, qiniu)
}

func (SettingsApi) UpdateQiNiu(c *gin.Context) {
	var cr config.QiNiu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	// 只将这一部分修改
	global.Config.QiNiu = cr
	if !core.SetYaml() {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "")
}
