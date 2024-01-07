package routers

import (
	"github.com/gin-gonic/gin"
	v1 "wild_goose_gin/api/v1"
	"wild_goose_gin/global"
	"wild_goose_gin/middleware/jwt"
	"wild_goose_gin/models"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	// 将静态文件目录设置为 "uploads" 目录
	// 这会将 "localhost:4040/api/static" 映射到 "./uploads" 目录下的文件
	r.Static("/api/static", "./uploads")

	r.POST("/api/login", v1.ApiGroupApp.UserApi.Login)
	// 需要登陆
	r.Use(jwt.JwtAuth())
	// 创建ws全局的连接
	//r.GET("ws_conn", service.AppService.WebsocketService.WebsocketConn)
	// 聊天服务
	//r.GET("ws_chat", service.AppService.WebsocketService.WebsocketChat)
	apiGroup := r.Group("api")
	// 系统设置api
	SettingsRouter(apiGroup)
	ComponentRouter(apiGroup)
	UserRouter(apiGroup)
	TaskRouter(apiGroup)
	ManualRouter(apiGroup)
	ModifyRouter(apiGroup)
	CustomRouter(apiGroup)
	MenuRouter(apiGroup)
	RouteRouter(apiGroup)
	RoleRouter(apiGroup)

	//将全部路由信息存入数据库
	Permissions := []models.Permission{}
	routes := r.Routes()
	for _, route := range routes {
		permission := models.Permission{
			Method: route.Method,
			Path:   route.Path,
		}
		Permissions = append(Permissions, permission)
	}
	var permissionModel models.Permission
	if err := permissionModel.DeleteAllRecords(); err != nil {
		global.Logrus.Fatal("清空权限失败")
	}
	if err := permissionModel.AddRoutes(Permissions); err != nil {
		global.Logrus.Fatal("添加权限失败")
	}
	return r
}
