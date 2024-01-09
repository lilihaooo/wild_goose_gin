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
	Node            common_type.TaskNodeType
	Demand          common_type.TaskDemandType
	Share           common_type.TaskShareType
	Certificates    *[]Certificate `gorm:"many2many:task_certificate;"`
	Modifies        []Modify       `gorm:"many2many:task_modify;"`
	GroupID         uint           `gorm:"comment:分组ID"`
	PlanReleaseDate time.Time      `gorm:"comment:计划放行日期"`
	Remark          string         `gorm:"comment:备注"`
	UserID          *uint          `gorm:"comment:工作者"`
	User            *User
}

func (t *Task) GetAllRecord(conditions string) (list []Task, err error) {
	err = global.DB.Model(t).
		Where(conditions).
		Preload("Component").
		Preload("Custom").
		Preload("Certificates").
		Preload("Modifies").
		Preload("User").
		Find(&list).Error
	return
}

func (t *Task) GetAllPagingRecord(limit int, offset int, conditions string) (list []Task, count int64, err error) {
	q := global.DB.Model(t).
		Where(conditions)
	err = q.Count(&count).Error
	err = q.Limit(limit).
		Offset(offset).
		Preload("Component").
		Preload("Custom").
		Preload("Certificates").
		Preload("Modifies").
		Preload("User").
		Find(&list).Error
	return
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

func (t *Task) TakeOneRecordByID() (*Task, error) {
	err := global.DB.Model(t).Preload("Certificates").Preload("Component.Manual").Take(&t, t.ID).Error
	return t, err
}

func (t *Task) TakeOneRecordByID2() (*Task, error) {
	err := global.DB.Model(t).Preload("Custom").Preload("Modifies").Preload("Component.Manual").Take(&t, t.ID).Error
	return t, err
}
