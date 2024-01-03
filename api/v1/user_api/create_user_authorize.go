package user_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils"
)

type CreateUserAuthorizeReq struct {
	UserID         uint   `json:"user_id" validate:"required" label:"用户ID"`
	CertificateIDs []uint `json:"certificate_ids" validate:"required" label:"证书IDs"`
	ManualID       uint   `json:"manual_id" validate:"required" label:"手册ID"`
}

func (UserApi) CreateUserAuthorize(c *gin.Context) {
	var req CreateUserAuthorizeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	// 验证用户是否存在
	var userModel models.User
	userModel.ID = req.UserID
	user, err := userModel.TakeOneUser()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	//验证手册是否存在
	var manualModel models.Manual
	manualModel.ID = req.ManualID
	manual, err := manualModel.TakeOneManual()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	//验证证书是否存在  (过滤)
	var certificateModel models.Certificate
	certificates, err := certificateModel.GetCertificatesByIDs(req.CertificateIDs)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	if len(certificates) == 0 {
		response.FailWithMsg(c, response.INVALID_PARAMS, "证书id不能为空")
		return
	}

	err = service.AppService.CreateUserAuthorize(user.ID, manual.ID, certificates)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error())
		return
	}
	response.OkWithMsg(c, "添加成功")
}
