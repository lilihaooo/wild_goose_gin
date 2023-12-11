package component_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type modifyInfo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ModifyNum string `json:"modify_num"`
	Version   string `json:"version"`
}

func (ComponentApi) GetComponentModifyList(c *gin.Context) {
	IDStr := c.Query("id")
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	ID, _ := strconv.Atoi(IDStr)

	var modify models.Modify
	list, err := modify.GetListByComponentID(uint(ID))
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}

	modifyInfoList := []modifyInfo{}
	if len(list) > 0 {
		for _, item := range list {
			one := modifyInfo{
				ID:        item.ID,
				Title:     item.Title,
				ModifyNum: item.Num,
				Version:   item.Version,
			}
			modifyInfoList = append(modifyInfoList, one)
		}
	}

	response.OkWithData(c, modifyInfoList)
}
