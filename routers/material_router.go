package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func MaterialRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("material/search/list", v1.ApiGroupApp.MaterialApi.GetMaterialSearchList)
}
