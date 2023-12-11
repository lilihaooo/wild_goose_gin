package models

import "wild_goose_gin/global"

type Certificate struct {
	ID    uint   `json:"id"`
	Title string `gorm:"comment:证书title" json:"title"`
}

func (c *Certificate) GetCertificatesByIDs(ids []uint) (certificates []Certificate, err error) {
	err = global.DB.Where("id IN ?", ids).Find(&certificates).Error
	return certificates, err
}
