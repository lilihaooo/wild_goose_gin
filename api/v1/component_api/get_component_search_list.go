package component_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type getComponentSearchListResp struct {
	ID   uint   `json:"id"`
	Pn   string `json:"pn"`
	Name string `json:"name"`
}

func (ComponentApi) GetComponentSearchList(c *gin.Context) {
	groupIDStr := c.Query("group_id")
	if groupIDStr == "" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "group_id不能为空")
		return
	}
	groupID, _ := strconv.Atoi(groupIDStr)
	var component models.Component
	list, err := component.GetComponentsByGroupID(groupID)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	respItem := getComponentSearchListResp{}
	respList := []getComponentSearchListResp{}
	for _, item := range list {
		respItem.ID = item.ID
		respItem.Pn = item.PN
		respItem.Name = item.Name
		respList = append(respList, respItem)
	}
	response.OkWithData(c, respList)
	return
}
