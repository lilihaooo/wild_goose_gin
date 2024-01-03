package component_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type certificateInfo struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (ComponentApi) GetComponentCertificateList(c *gin.Context) {
	IDStr := c.Query("id")
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, "id不能为空")
		return
	}
	ID, _ := strconv.Atoi(IDStr)
	var model models.Component
	model.ID = uint(ID)
	component, err := model.GetComponentCertificateByID()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}

	certificateInfoList := []certificateInfo{}
	if len(component.Certificates) > 0 {
		for _, item := range component.Certificates {
			one := certificateInfo{
				ID:    item.ID,
				Title: item.Title,
			}
			certificateInfoList = append(certificateInfoList, one)
		}
	}

	response.OkWithData(c, certificateInfoList)
}
