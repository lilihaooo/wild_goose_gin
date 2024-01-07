package models

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models/common_type"
)

type Menu struct {
	ID       *uint                `json:"id"`
	Icon     string               `json:"icon"`
	Path     string               `json:"path"`
	Title    string               `json:"title"`
	ParentID *uint                `json:"parent_id,omitempty"`
	Type     common_type.MenuType `json:"type"`
	Sort     int32                `json:"sort"`
	Subs     []*Menu              `json:"subs" gorm:"foreignKey:ParentID"`
}

func (m *Menu) GetAllRecord() (menus []*Menu, err error) {
	err = global.DB.Order("sort desc").Find(&menus).Error
	return
}

func (m *Menu) GetOneRecordById() (*Menu, error) {
	err := global.DB.Take(&m, m.ID).Error
	return m, err
}

func (m *Menu) SaveMenu() error {
	return global.DB.Save(m).Error
}
