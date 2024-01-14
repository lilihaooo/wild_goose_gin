package task_api

import (
	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils"
)

type CreateTaskRequest struct {
	PN              string                     `json:"pn" validate:"required" label:"件号"`
	SN              string                     `json:"sn" validate:"required" label:"序号"`
	CustomID        uint                       `json:"custom_id" validate:"required" label:"客户id"`
	DemandType      common_type.TaskDemandType `json:"demand_type" validate:"required" label:"任务要求"`
	CertificateIDs  *[]uint                    `json:"certificate_ids"`
	ModifyIDs       []uint                     `json:"modify_ids"`
	GroupID         uint                       `json:"group_id" validate:"required" label:"分组id"`
	PlanReleaseDate string                     `json:"plan_release_date" validate:"required" label:"计划放行日期"`
	Remark          string                     `json:"remark" validate:"max=5" label:"备注"`
}

func (TaskApi) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "接收参数失败"+err.Error()) // todo err
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	// 获得附件信息
	componentModel := models.Component{}
	componentModel.PN = req.PN
	componentModel.GroupID = req.GroupID
	component, err := componentModel.GetComponentByPN()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	if component.ID == 0 {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "附件信息不存在")
		return
	}
	if component.State == 0 {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "附件状态异常")
		return
	}
	// 获得客户信息
	customModel := models.Custom{}
	customModel.ID = req.CustomID
	custom, err := customModel.GetRecordByID()
	if err != nil {
		response.FailWithMsg(c, response.ERROR_EXIST_RECODE, "")
		return
	}

	// 根据改装ids, 获得改装
	modifyModel := models.Modify{
		ComponentID: component.ID,
	}
	modifies, err := modifyModel.GetModifiesByIDs(req.ModifyIDs)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "获取改装失败"+err.Error()) // todo err
		return
	}
	if len(modifies) != len(req.ModifyIDs) {
		response.FailWithMsg(c, response.FAIL_OPER, "改装有误")
		return
	}

	// 根据证书ids, 获得证书信息
	certificateModel := models.Certificate{}
	certificates := []models.Certificate{}
	if req.CertificateIDs != nil {
		certificates, err = certificateModel.GetCertificatesByIDs(*req.CertificateIDs)
		if err != nil {
			response.FailWithMsg(c, response.FAIL_OPER, "获取证书失败"+err.Error()) // todo err
			return
		}
		if len(certificates) != len(*req.CertificateIDs) {
			response.FailWithMsg(c, response.FAIL_OPER, "证书有误")
			return
		}
	}

	// 解析时间字符串
	parsedTime, err := dateparse.ParseAny(req.PlanReleaseDate)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "解析日期失败"+err.Error()) // todo err
		return
	}

	// 生成任务号
	TaskService := service.AppService.TaskService
	taskNum, err := TaskService.GenerateTaskNum(component.Group)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "生成任务号失败"+err.Error()) // todo err
		return
	}
	// 检查任务号是否存在(包括已经软删除的)
	var task models.Task
	taskNumIsExist := task.TaskNumIsExist(taskNum)
	if taskNumIsExist {
		response.FailWithMsg(c, response.FAIL_OPER, "任务号已存在")
		return
	}

	task.ComponentID = component.ID
	task.SN = req.SN
	task.CustomID = custom.ID
	task.Node = common_type.TaskStart
	task.Share = common_type.UnShared
	task.Demand = req.DemandType
	task.Modifies = modifies
	task.GroupID = component.GroupID
	task.Certificates = &certificates
	task.Remark = req.Remark
	task.PlanReleaseDate = parsedTime
	task.TaskNum = taskNum

	var isClaim = false
	if req.DemandType == common_type.Claim {
		isClaim = true
	}

	// 任务下发
	err = service.AppService.TaskService.CreateTask(&task, isClaim)
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, err.Error()) // todo 错误
		return
	}
	response.OkWithMsg(c, "添加成功")
}
