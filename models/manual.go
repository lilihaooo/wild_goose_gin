package models

import (
	"wild_goose_gin/global"
)

type Manual struct {
	ID        uint        `json:"id"`
	Num       string      `gorm:"comment:件名" json:"name"`
	Version   string      `gorm:"comment:版本" json:"version"`
	Users     *[]User     `gorm:"many2many:user_manual;"`
	Materials *[]Material `gorm:"many2many:manual_material;"`
}

func (m *Manual) GetList() (list []Manual, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}

func (m *Manual) TakeOneManual() (*Manual, error) {
	err := global.DB.Take(m).Error
	return m, err
}
