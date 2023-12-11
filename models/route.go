package models

import (
	"wild_goose_gin/global"
)

type Route struct {
	ID     uint   `json:"id"`
	Method string `gorm:"comment:方法" json:"method"`
	Path   string `gorm:"comment:路径" json:"path"`
}

func (r *Route) GetRouteList() (list []Route, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}

func (r *Route) DeleteAllRecords() error {
	return global.DB.Exec("DELETE FROM route").Error
}

func (r *Route) AddRoutes(routes []Route) error {
	return global.DB.Create(&routes).Error
}
