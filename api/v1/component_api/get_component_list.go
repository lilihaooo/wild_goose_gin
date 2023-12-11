package component_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	serializer "wild_goose_gin/serialize"
)

func (ComponentApi) GetComponentList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	var model models.Component
	list, count, err := model.GetAllRecord(req)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	resList := serializer.BuildComponents(list)
	response.OkWithList(c, resList, count)
}
