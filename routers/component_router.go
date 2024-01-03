package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func ComponentRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("component", v1.ApiGroupApp.ComponentApi.CreateComponent)
	v1Group.PUT("component", v1.ApiGroupApp.ComponentApi.UpdateComponent)
	v1Group.POST("component/all", v1.ApiGroupApp.ComponentApi.GetComponentList)
	v1Group.POST("component/del", v1.ApiGroupApp.ComponentApi.DeleteComponents)
	v1Group.GET("component/search/list", v1.ApiGroupApp.ComponentApi.GetComponentSearchList)
	v1Group.GET("component/change/state", v1.ApiGroupApp.ComponentApi.ComponentChangeState)
	v1Group.GET("component/info", v1.ApiGroupApp.ComponentApi.GetComponentInfo)
	v1Group.GET("component/modify/list", v1.ApiGroupApp.ComponentApi.GetComponentModifyList)
	v1Group.GET("component/certificate/list", v1.ApiGroupApp.ComponentApi.GetComponentCertificateList)
}
