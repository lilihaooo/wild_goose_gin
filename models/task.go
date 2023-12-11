package models

import (
	"time"
	"wild_goose_gin/global"
	"wild_goose_gin/models/common_type"
)

type Task struct {
	Model
	TaskNum         string `gorm:"comment:任务号"`
	ComponentID     uint   `gorm:"comment:附件ID"`
	Component       Component
	SN              string `gorm:"comment:附件序号"`
	CustomID        uint   `gorm:"comment:客户ID"`
	Custom          Custom
	NodeType        common_type.TaskNodeType
	DemandType      common_type.TaskDemandType
	Certificates    []Certificate `gorm:"many2many:task_certificate;"`
	Modifies        []Modify      `gorm:"many2many:task_modify;"`
	GroupID         uint          `gorm:"comment:分组ID" json:"group_id"`
	PlanReleaseDate time.Time     `gorm:"comment:计划放行日期" json:"plan_release_date"`
	Remark          string        `gorm:"comment:备注" json:"remark"`
}

func (t *Task) GetAllRecord(offset int, limit int) (list []Task, count int64, err error) {
	query := global.DB.Model(t)
	query.Count(&count)
	err = query.Offset(offset).Limit(limit).Find(&list).Error
	return list, count, err
}

func (t *Task) GetGroupThisMonthCount(groupID uint) (count int64, err error) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	// 查询本年本月添加的最后一条记录的 ID
	err = global.DB.Model(t).Unscoped().Where("group_id = ?", groupID).Where("created_at BETWEEN ? AND ?", firstOfMonth, lastOfMonth).
		Count(&count).Error
	return
}

func (t *Task) TaskNumIsExist(taskNum string) bool {
	err := global.DB.Where("task_num = ?", taskNum).Take(t).Error
	if err != nil {
		return false
	}
	return true
}

func (t *Task) AddOneRecord() error {
	return global.DB.Create(t).Error
}
