package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func ManualRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.GET("manual/list", v1.ApiGroupApp.ManualApi.GetManualList)
	v1Group.POST("manual/material/list", v1.ApiGroupApp.ManualApi.GetManualMaterialList)
}
