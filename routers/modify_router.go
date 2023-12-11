package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func ModifyRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("modify/create", v1.ApiGroupApp.ModifyApi.CreateModify)
	v1Group.DELETE("modify/del", v1.ApiGroupApp.ModifyApi.DeleteModify)
}
