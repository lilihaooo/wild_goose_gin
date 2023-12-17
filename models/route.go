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
func (r *Route) GetRouteListByIDs(ids []uint) (list []*Route, err error) {
	err = global.DB.Where("id IN (?)", ids).Find(&list).Error
	return list, err
}

func (r *Route) DeleteAllRecords() error {
	tx := global.DB.Begin()
	err := tx.Exec(" DELETE FROM menu_route  WHERE route_id IN (SELECT id FROM route)").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Exec("DELETE FROM route").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Route) AddRoutes(routes []Route) error {
	return global.DB.Create(&routes).Error
}
