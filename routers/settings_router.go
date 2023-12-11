package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func SettingsRouter(appGroup *gin.RouterGroup) {
	// v1SettingsApi 接口
	v1SettingsGroup := appGroup.Group("v1/settings")

	v1SettingsGroup.GET("site_info", v1.ApiGroupApp.SettingsApi.GetSiteInfo)
	v1SettingsGroup.PUT("site_info", v1.ApiGroupApp.SettingsApi.UpdateSiteInfo)

	v1SettingsGroup.GET("email", v1.ApiGroupApp.SettingsApi.GetEmail)
	v1SettingsGroup.PUT("email", v1.ApiGroupApp.SettingsApi.UpdateEmail)
	v1SettingsGroup.GET("jwt", v1.ApiGroupApp.SettingsApi.GetJwt)
	v1SettingsGroup.PUT("jwt", v1.ApiGroupApp.SettingsApi.UpdateJwt)

	v1SettingsGroup.GET("qiniu", v1.ApiGroupApp.SettingsApi.GetQiNiu)
	v1SettingsGroup.PUT("qiniu", v1.ApiGroupApp.SettingsApi.UpdateQiNiu)
}
