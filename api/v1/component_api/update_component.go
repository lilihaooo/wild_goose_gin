package component_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type UpdateComponentRequest struct {
	ID             uint   `json:"id" validate:"required" label:"ID"`
	CertificateIDs []uint `json:"certificate_ids"`
}

func (ComponentApi) UpdateComponent(c *gin.Context) {
	var req UpdateComponentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var componentModel models.Component
	// 获取附件信息
	componentModel.ID = req.ID
	component, err := componentModel.GetComponentByID()
	if err != nil {
		response.FailWithMsg(c, response.ERROR_NOT_EXIST_RECODE, err.Error()) // todo err
		return
	}
	// 根据证书ids, 获得证书信息
	certificateModel := models.Certificate{}
	certificates, err := certificateModel.GetCertificatesByIDs(req.CertificateIDs)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "获取证书失败"+err.Error()) // todo err
		return
	}
	if len(certificates) != len(req.CertificateIDs) {
		response.FailWithMsg(c, response.FAIL_OPER, "证书有误")
		return
	}
	component.Certificates = certificates

	if err := component.UpdateRecordAndAssociation(); err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error()) // todo error
		return
	}
	response.OkWithMsg(c, "修改成功")
}
