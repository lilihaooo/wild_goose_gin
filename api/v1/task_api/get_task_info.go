package task_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
)

type taskInfoResp struct {
	ID            uint                       `json:"id"`
	TaskNum       string                     `json:"task_num"`
	ManualNum     string                     `json:"manual_num"`
	ManualID      uint                       `json:"manual_id"`
	ComponentName string                     `json:"component_name"`
	CustomName    string                     `json:"custom_name"`
	Demand        common_type.TaskDemandType `json:"demand"`
	ModifyNums    []string                   `json:"modify_nums"`
}

func (TaskApi) GetTaskInfo(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	id, _ := strconv.Atoi(idStr)

	var model models.Task
	model.ID = uint(id)

	task, err := model.TakeOneRecordByID2()
	if err != nil {
		if err != nil {
			global.Logrus.Error(err)
			response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
			return
		}
	}

	var modifyNums []string
	if task.Modifies != nil {
		for _, one := range task.Modifies {
			modifyNums = append(modifyNums, one.Num)
		}
	}

	resp := taskInfoResp{
		ID:            task.ID,
		ManualNum:     task.Component.Manual.Num,
		ManualID:      task.Component.ManualID,
		ComponentName: task.Component.Name,
		CustomName:    task.Custom.Name,
		Demand:        task.Demand,
		ModifyNums:    modifyNums,
	}

	response.OkWithData(c, resp)
}
