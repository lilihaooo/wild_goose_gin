package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func CustomRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.GET("custom/input/list", v1.ApiGroupApp.CustomApi.GetCustomInputList)
}
