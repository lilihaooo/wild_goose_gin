package models

import (
	"wild_goose_gin/global"
)

type Manual struct {
	ID      uint   `json:"id"`
	Num     string `gorm:"comment:件名" json:"name"`
	Version string `gorm:"comment:版本" json:"version"`
}

func (m *Manual) GetInputList() (list []Manual, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}
