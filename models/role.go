package models

import (
	"wild_goose_gin/global"
)

type Role struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Users       []User       `gorm:"many2many:user_role;" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permission;" json:"permissions,omitempty"`
	Menus       []*Menu      `gorm:"many2many:role_menu;" json:"menus,omitempty"`
}

func (r *Role) GetAllRoleList() (list []Role, err error) {
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

func (c *Role) GetRolesByIDs(ids []uint) (roles []Role, err error) {
	err = global.DB.Where("id IN ?", ids).Find(&roles).Error
	return roles, err
}
