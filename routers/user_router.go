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
	v1Group.POST("user/update", v1.ApiGroupApp.UserApi.UpdateUser)
	v1Group.DELETE("user/del", v1.ApiGroupApp.UserApi.DeleteUser)
	v1Group.GET("user/info", v1.ApiGroupApp.UserApi.UserInfo)
	v1Group.GET("user/all/select/list", v1.ApiGroupApp.UserApi.UserAllSelectList)
	v1Group.POST("user/authorize/list", v1.ApiGroupApp.UserApi.UserAuthorizeUserList)                                  // 用户授权页面的user_list
	v1Group.GET("user/authorize/list", v1.ApiGroupApp.UserApi.GetUserAuthorizeList)                                    // 用户的授权list
	v1Group.POST("user/authorize", v1.ApiGroupApp.UserApi.CreateUserAuthorize)                                         // 添加用户授权
	v1Group.POST("user/user_manual_certificate/state/change", v1.ApiGroupApp.UserApi.ChangeUserManualCertificateState) // 用户的授权证书状态改变
	v1Group.GET("user/task/optional/list", v1.ApiGroupApp.UserApi.UserTaskOptionalList)                                // 任务可选择的用户列表
	v1Group.GET("user/role/list", v1.ApiGroupApp.UserApi.GetUserRoleIDs)                                               // 获得user.role列表
	v1Group.POST("user/role", v1.ApiGroupApp.UserApi.SetUserRole)                                                      // 获得user.role列表
}
