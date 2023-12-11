package routers

import (
	"github.com/gin-gonic/gin"
	v1 "wild_goose_gin/api/v1"
)

func UserRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("user/set_avatar", v1.ApiGroupApp.UserApi.SetAvatar)
	v1Group.POST("user/list", v1.ApiGroupApp.UserApi.UserList)
	v1Group.POST("user/create", v1.ApiGroupApp.UserApi.CreateUser)
	v1Group.DELETE("user/del", v1.ApiGroupApp.UserApi.DeleteUser)
	v1Group.GET("user/info", v1.ApiGroupApp.UserApi.UserInfo)
}
