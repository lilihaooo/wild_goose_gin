package manual_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/serialize/manual_serialize"
)

func (ManualApi) GetManualList(c *gin.Context) {
	var model models.Manual
	list, err := model.GetList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := manual_serialize.BuildManuals(list)
	response.OkWithData(c, resList)
}
