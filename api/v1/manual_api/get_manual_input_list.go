package manual_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type ManualInput struct {
	ManualID uint   `json:"manual_id"`
	Num      string `json:"manual_num"`
}

func (ManualApi) GetManualInputList(c *gin.Context) {
	var model models.Manual
	list, err := model.GetInputList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := ManualInputList(list)
	response.OkWithData(c, resList)
}

func BuildManualInputListRes(item models.Manual) ManualInput {
	return ManualInput{
		ManualID: item.ID,
		Num:      item.Num,
	}
}

func ManualInputList(items []models.Manual) (manuals []ManualInput) {
	for _, item := range items {
		manual := BuildManualInputListRes(item)
		manuals = append(manuals, manual)
	}
	return manuals
}
