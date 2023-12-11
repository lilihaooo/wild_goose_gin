package models

import "wild_goose_gin/global"

type Menu struct {
	ID       uint     `json:"id"`
	Icon     string   `json:"icon"`
	Path     string   `json:"path"`
	Title    string   `json:"title"`
	ParentID uint     `json:"parent_id"`
	Routes   []*Route `gorm:"many2many:menu_route;"`
	Subs     []*Menu  `json:"subs" gorm:"foreignKey:ParentID"`
}

func (m *Menu) GetAllRecord() (menus []*Menu, err error) {
	err = global.DB.Find(&menus).Error
	return
}
