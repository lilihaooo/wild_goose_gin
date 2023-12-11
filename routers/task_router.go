package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func TaskRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("task/all", v1.ApiGroupApp.TaskApi.GetTaskList)
	v1Group.POST("task/add", v1.ApiGroupApp.TaskApi.CreateTask)
}
