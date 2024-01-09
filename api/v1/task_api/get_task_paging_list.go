package task_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/serialize"
)

func (TaskApi) GetTaskPagingList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, err.Error()) // todo 去掉真实的error
		return
	}
	conditions := []request.Condition{}
	if req.Conditions != nil {
		for _, one := range *req.Conditions {
			if one.Key != "share" || one.Value != "0" {
				conditions = append(conditions, one)
			}
		}
	}
	conditionStr := request.GetConditionStr(conditions)
	var model models.Task
	list, count, err := model.GetAllPagingRecord(req.PageSize, req.GetOffset(), conditionStr)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := serializer.BuildTasks(list)
	response.OkWithList(c, resList, count)
}