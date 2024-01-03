package user_ser

import (
	"errors"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
)

func (s UserService) CreateUserAuthorize(userID, manualID uint, certificates []models.Certificate) error {
	var umModel models.UserManual
	umModel.UserID = userID
	umModel.ManualID = manualID
	um, err := umModel.FindOneByUserIDAndManualID()
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	umc := models.UserManualCertificate{}
	EXcids := []uint{}
	if um.ID != 0 {
		// 用户存在手册
		umc.UserManualID = um.ID
		EXcids, err = umc.FindCertificateIDsByUserManualID() // 查询用户手册存在的证书IDs
		if err != nil {
			global.Logrus.Error(err)
			return err
		}
		var newUMC []models.UserManualCertificate
		if len(EXcids) > 0 {
			cMap := make(map[uint]uint)
			for _, one := range EXcids {
				cMap[one] = one
			}
			for _, item := range certificates {
				if cMap[item.ID] == 0 {
					newUMC = append(newUMC, models.UserManualCertificate{
						UserManualID:  um.ID,
						CertificateID: item.ID,
						State:         common_type.AuthorizeStageNormal,
					})
				}
			}
			if len(newUMC) > 0 {
				err = global.DB.Create(&newUMC).Error
				return err
			}
			err = errors.New("没有数据新增")
			return err
		} else {
			var newUMC []models.UserManualCertificate
			for _, item := range certificates {
				newUMC = append(newUMC, models.UserManualCertificate{
					UserManualID:  um.ID,
					CertificateID: item.ID,
					State:         common_type.AuthorizeStageNormal,
				})
			}
			err = global.DB.Create(&newUMC).Error
			return err
		}
	} else {
		// 创建全新的关联数据
		umModel.UserID = userID
		umModel.ManualID = manualID
		tx := global.DB.Begin()
		err = tx.Create(&umModel).Error
		if err != nil {
			global.Logrus.Error(err)
			tx.Rollback()
			return err
		}
		var newUMC []models.UserManualCertificate
		for _, item := range certificates {
			newUMC = append(newUMC, models.UserManualCertificate{
				UserManualID:  umModel.ID,
				CertificateID: item.ID,
				State:         common_type.AuthorizeStageNormal,
			})
		}
		err = tx.Create(&newUMC).Error
		if err != nil {
			global.Logrus.Error(err)
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return err
}
