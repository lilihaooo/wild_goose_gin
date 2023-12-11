package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func RouteRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.GET("route/all/list", v1.ApiGroupApp.RouteApi.GetAllRouteList)
}
