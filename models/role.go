package models

import (
	"wild_goose_gin/global"
)

type Role struct {
	ID    uint
	Title string
	Menus []*Menu `gorm:"many2many:role_menu;"`
}

func (r *Role) GetRoleList() (list []Role, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}

func (r *Role) AddRole() error {
	return global.DB.Create(&r).Error
}

func (r *Role) GetRoleMenuList() (menus []*Menu, err error) {
	err = global.DB.Preload("Menus").Find(r).Error
	return r.Menus, err
}
