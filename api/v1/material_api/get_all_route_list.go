package material_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/serialize/material_serialize"
)

func (MaterialApi) GetMaterialSearchList(c *gin.Context) {
	req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	var model models.Material
	list, count, err := model.GetAllRecord(req)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	res := material_serialize.BuildMaterials(list)

	response.OkWithList(c, res, count)
}
