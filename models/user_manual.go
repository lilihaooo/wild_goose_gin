package models

import (
	"wild_goose_gin/global"
)

type UserManual struct {
	ID                     uint `gorm:"primaryKey"`
	UserID                 uint `gorm:"foreignKey:UserID;references:ID"`
	User                   User
	ManualID               uint `gorm:"foreignKey:ManualID;references:ID"`
	Manual                 Manual
	Certificates           *[]Certificate `gorm:"many2many:user_manual_certificate;"` // 仅用于创建表关系, 获取证书是不使用该属性, 因为需要获得证书的状态, 使用UserManualCertificates
	UserManualCertificates *[]UserManualCertificate
}

func (u *UserManual) GetRecordByUserID() (userManuals []UserManual, err error) {
	err = global.DB.Preload("UserManualCertificates.Certificate").Preload("Manual").Where("user_id = ?", u.UserID).Find(&userManuals).Error
	return
}

func (u *UserManual) GetRecordsByManualID() (ums []UserManual, err error) {
	err = global.DB.Preload("User").Where("manual_id = ?", u.ManualID).Find(&ums).Error
	return
}

func (u *UserManual) GetRecordsByManualIDWithCertificate() (ums []UserManual, err error) {
	err = global.DB.Preload("User").Preload("UserManualCertificates.Certificate").Where("manual_id = ?", u.ManualID).Find(&ums).Error
	return
}

func (u *UserManual) FindOneByUserIDAndManualID() (userManual *UserManual, err error) {
	err = global.DB.Where("user_id = ? And manual_id = ?", u.UserID, u.ManualID).Find(&userManual).Error
	return userManual, err
}
