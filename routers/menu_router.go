package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func MenuRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.GET("menu/all/tree", v1.ApiGroupApp.MenuApi.GetAllMenuTree)
	v1Group.GET("menu/role/tree", v1.ApiGroupApp.MenuApi.GetRoleMenuTree)
	v1Group.POST("menu/add", v1.ApiGroupApp.MenuApi.CreateMenu)
	v1Group.PUT("menu", v1.ApiGroupApp.MenuApi.UpdateMenu)
}
