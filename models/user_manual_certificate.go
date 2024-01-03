package models

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models/common_type"
)

type UserManualCertificate struct {
	Model
	UserManualID  uint `gorm:"foreignKey:UserManualID;references:ID"`
	UserManual    UserManual
	CertificateID uint `gorm:"foreignKey:CertificateID;references:ID"`
	Certificate   Certificate
	State         common_type.AuthorizeStageType `gorm:"comment:状态" json:"state"`
}

func (u *UserManualCertificate) FindCertificateIDsByUserManualID() (certificateIDs []uint, err error) {
	err = global.DB.Model(u).Select("certificate_id").Where("user_manual_id = ?", u.UserManualID).Scan(&certificateIDs).Error
	return
}

func (u *UserManualCertificate) ChangeState() error {
	err := global.DB.Find(u).Error
	if err != nil {
		return err
	}
	if u.State == common_type.AuthorizeStageNormal {
		u.State = common_type.AuthorizeStageStop
	} else {
		u.State = common_type.AuthorizeStageNormal
	}
	err = global.DB.Save(u).Error
	return err
}
