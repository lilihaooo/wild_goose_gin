package routers

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/api/v1"
)

func TaskRouter(appGroup *gin.RouterGroup) {
	v1Group := appGroup.Group("v1")
	v1Group.POST("task/all", v1.ApiGroupApp.TaskApi.GetTaskAllList)
	v1Group.POST("task/all/paging", v1.ApiGroupApp.TaskApi.GetTaskPagingList)
	v1Group.POST("task/my_tasking/paging", v1.ApiGroupApp.TaskApi.GetMyTaskingPagingList)
	v1Group.POST("task/add", v1.ApiGroupApp.TaskApi.CreateTask)
	v1Group.POST("task/share", v1.ApiGroupApp.TaskApi.TaskShare)
	v1Group.GET("task", v1.ApiGroupApp.TaskApi.GetTaskInfo)
}
