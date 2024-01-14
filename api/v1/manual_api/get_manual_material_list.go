package manual_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/request"
	"wild_goose_gin/pkg/response"
	manual_serialize "wild_goose_gin/serialize/manual_material_serialize"
	"wild_goose_gin/utils"
)

type GetManualMaterialListReq struct {
	PReq     request.PaginationReq
	ManualID uint `json:"manual_id" validate:"required"`
}

func (ManualApi) GetManualMaterialList(c *gin.Context) {
	//manualIDStr := c.Query("manual_id")
	//if manualIDStr == "" {
	//	response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
	//	return
	//}
	//manualID, _ := strconv.Atoi(manualIDStr)

	req := GetManualMaterialListReq{
		PReq: *request.NewPaginationReq(),
	}

	//req := request.NewPaginationReq()
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}

	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}

	var model models.ManualMaterial
	model.ManualID = req.ManualID
	// 查询
	manualMaterial, count, err := model.GetListByManualID(&req.PReq)
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.FAIL_OPER, "查询失败")
		return
	}
	// 序列化
	res := manual_serialize.BuildManualMaterials(manualMaterial)
	response.OkWithList(c, res, count)
}
