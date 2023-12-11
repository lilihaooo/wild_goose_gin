package custom_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type CustomInput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (CustomApi) GetCustomInputList(c *gin.Context) {
	var model models.Custom
	list, err := model.GetInputList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := CustomInputList(list)
	response.OkWithData(c, resList)
}

func BuildCustomInputListRes(item models.Custom) CustomInput {
	return CustomInput{
		ID:   item.ID,
		Name: item.Name,
	}
}

func CustomInputList(items []models.Custom) (customs []CustomInput) {
	for _, item := range items {
		manual := BuildCustomInputListRes(item)
		customs = append(customs, manual)
	}
	return customs
}
