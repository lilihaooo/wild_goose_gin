package component_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type componentInfo struct {
	Name          string                         `json:"name"`
	PN            string                         `json:"pn"`
	ManualNum     string                         `json:"manual_num"`
	Group         string                         `json:"group"`
	ModifiesCount int                            `json:"modifies_count"`
	CreatedAt     string                         `json:"created_at"`
	State         common_type.ComponentStageType `json:"state"`
	IncomeTotal   uint                           `json:"income_total"`
}

func (ComponentApi) GetComponentInfo(c *gin.Context) {
	IDStr := c.Query("id")
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	ID, _ := strconv.Atoi(IDStr)
	var component models.Component
	component.Model.ID = uint(ID)
	data, err := component.GetDetailInfoByID()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}

	res := componentInfo{
		Name:          data.Name,
		PN:            data.PN,
		ManualNum:     data.Manual.Num,
		Group:         data.Group.Name,
		ModifiesCount: len(data.Modifies),
		CreatedAt:     utils.TimeFormat(data.CreatedAt),
		State:         data.State,
		IncomeTotal:   data.IncomeTotal,
	}
	response.OkWithData(c, res)
}
