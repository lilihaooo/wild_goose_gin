package models

import "wild_goose_gin/global"

type Custom struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"comment:件名" json:"name"`
	Address string `gorm:"comment:地址" json:"address"`
	Email   string `gorm:"comment:邮箱" json:"email"`
	Tel     string `gorm:"comment:手机号" json:"tel"`
}

func (c *Custom) GetInputList() (list []Custom, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}

func (c *Custom) GetRecordByID() (*Custom, error) {
	err := global.DB.Where("id = ?", c.ID).Take(&c).Error
	return c, err
}
