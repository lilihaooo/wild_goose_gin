package task_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	serializer "wild_goose_gin/serialize"
)

func (TaskApi) GetTaskList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	offset := req.GetOffset()
	limit := req.PageSize
	var model models.Task
	list, count, err := model.GetAllRecord(offset, limit)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := serializer.BuildTasks(list)
	response.OkWithList(c, resList, count)
}
