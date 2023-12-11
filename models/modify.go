package models

import "wild_goose_gin/global"

type Modify struct {
	ID          uint   `json:"id"`
	Title       string `gorm:"comment:改装标题" json:"title"`
	Num         string `gorm:"comment:改装号" json:"num"`
	Version     string `gorm:"comment:版本" json:"version"`
	ComponentID uint   `gorm:"comment:部件ID" json:"component_id"`
}

func (m *Modify) GetListByComponentID(componentID uint) (list []Modify, err error) {
	err = global.DB.Where("component_id = ?", componentID).Find(&list).Error
	return list, err
}

func (m *Modify) AddOneRecord() error {
	return global.DB.Create(m).Error
}

func (m *Modify) DeleteOneRecord() error {
	return global.DB.Delete(m).Error
}

func (m *Modify) GetModifiesByIDs(ids []uint) (modifies []Modify, err error) {
	err = global.DB.Where("id IN ?", ids).Where("component_id = ?", m.ComponentID).Find(&modifies).Error
	return modifies, err
}
