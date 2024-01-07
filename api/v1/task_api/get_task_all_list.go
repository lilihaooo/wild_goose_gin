package task_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/serialize"
)

type GetTaskListReq struct {
	Conditions []request.Condition `json:"condition" validate:"required"`
}

func (TaskApi) GetTaskAllList(c *gin.Context) {
	var req GetTaskListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, err.Error()) // todo 去掉真实的error
		return
	}
	conditions := []request.Condition{}
	for _, one := range req.Conditions {
		if one.Key != "share" || one.Value != "0" {
			conditions = append(conditions, one)
		}
	}
	conditionStr := request.GetConditionStr(conditions)
	var model models.Task
	list, err := model.GetAllRecord(conditionStr)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := serializer.BuildTasks(list)
	response.OkWithData(c, resList)
}
