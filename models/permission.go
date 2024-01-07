package models

import (
	"wild_goose_gin/global"
)

type Permission struct {
	ID     uint   `json:"id"`
	Method string `gorm:"comment:方法" json:"method"`
	Path   string `gorm:"comment:路径" json:"path"`
	Roles  []Role `gorm:"many2many:role_permission;"`
}

func (r *Permission) GetRouteList() (list []Permission, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}
func (r *Permission) GetRouteListByIDs(ids []uint) (list []*Permission, err error) {
	err = global.DB.Where("id IN (?)", ids).Find(&list).Error
	return list, err
}

func (r *Permission) DeleteAllRecords() error {
	return global.DB.Exec("DELETE FROM permission").Error
}

func (r *Permission) AddRoutes(routes []Permission) error {
	return global.DB.Create(&routes).Error
}
