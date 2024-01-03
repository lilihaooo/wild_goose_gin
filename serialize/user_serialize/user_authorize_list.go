package user_serialize

import (
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"
)

type Authorize struct {
	ManualNum    string        `json:"manual_num"`
	Certificates []Certificate `json:"certificates"`
}

type Certificate struct {
	ID       uint                           `json:"id"`
	Name     string                         `json:"name"`
	CreateAt string                         `json:"created_at"`
	State    common_type.AuthorizeStageType `json:"state"`
}

func BuildUserAuthorize(item models.UserManual) Authorize {
	certificates := []Certificate{}
	if len(*item.UserManualCertificates) > 0 {
		for _, one := range *item.UserManualCertificates {
			certificates = append(certificates, Certificate{
				ID:       one.ID,
				Name:     one.Certificate.Title,
				State:    one.State,
				CreateAt: utils.TimeFormat_YMD(one.CreatedAt),
			})
		}
	}
	return Authorize{
		ManualNum:    item.Manual.Num,
		Certificates: certificates,
	}
}

func BuildUserAuthorizes(items []models.UserManual) (userManuals []Authorize) {
	for _, item := range items {
		user := BuildUserAuthorize(item)
		userManuals = append(userManuals, user)
	}
	return
}
