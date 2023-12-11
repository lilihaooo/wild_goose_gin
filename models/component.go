package models

import (
	"fmt"
	"wild_goose_gin/global"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/request"
)

type Component struct {
	Model
	Name        string `gorm:"comment:件名" json:"name"`
	PN          string `gorm:"comment:件号" json:"pn"`
	ManualID    uint   `gorm:"comment:手册ID" json:"manual_id"`
	Manual      Manual
	GroupID     uint `gorm:"comment:小组ID" json:"group_id"`
	Group       Group
	Modifies    []Modify                       `gorm:"foreignkey:ComponentID;constraint:OnDelete:CASCADE;" json:"modifies"`
	State       common_type.ComponentStageType `gorm:"comment:状态" json:"state"`
	IncomeTotal uint                           `gorm:"comment:入场数量" json:"income_total"`
}

func (c *Component) AddOneRecord() error {
	return global.DB.Create(c).Error
}

func (c *Component) GetAllRecord(req *request.PaginationReq) (list []Component, count int64, err error) {
	query := global.DB.Model(c)
	if req.Keyword != "" {
		query.Where("name like ?", "%"+req.Keyword+"%").Or("pn like ?", "%"+req.Keyword+"%")
	}
	query.Count(&count)
	if req.Sort != nil {
		orderStr := fmt.Sprintf("%s %s", req.Sort.Field, req.Sort.Order)
		query.Order(orderStr)
	}
	err = query.Preload("Group").Preload("Manual").Preload("Modifies").Offset(req.GetOffset()).Limit(req.PageSize).Find(&list).Error
	return list, count, err
}

func (c *Component) GetComponentsByGroupID(groupID int) (list []Component, err error) {
	err = global.DB.Model(c).Select("name, pn, id").Where("group_id = ?", groupID).Where("state = ?", common_type.ComponentStageNormal).Find(&list).Error
	return list, err
}

func (c *Component) ChangeState() error {
	err := global.DB.Find(c).Error
	if err != nil {
		return err
	}
	if c.State == common_type.ComponentStageNormal {
		c.State = common_type.ComponentStageStop
	} else {
		c.State = common_type.ComponentStageNormal
	}
	err = global.DB.Save(c).Error
	return err
}

func (c *Component) DeleteComponents(ids []int64) error {
	return global.DB.Preload("Modifies").Where("id IN ?", ids).Delete(c).Error
}

func (c *Component) GetDetailInfoByID() (*Component, error) {
	err := global.DB.Preload("Group").Preload("Manual").Preload("Modifies").Take(c).Error
	return c, err
}

func (c *Component) GetComponentByPN() (*Component, error) {
	err := global.DB.Model(c).Preload("Group").Where("pn = ?", c.PN).Find(&c).Error
	return c, err
}
