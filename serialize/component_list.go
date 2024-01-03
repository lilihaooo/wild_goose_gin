package serializer

import (
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"
)

type Component struct {
	ID             uint                           `json:"id"`
	Name           string                         `json:"name"`
	PN             string                         `json:"pn"`
	ManualNum      string                         `json:"manual_num"`
	ModifyCount    int                            `json:"modify_count"`
	Certificates   string                         `json:"certificates"`
	CertificateIDs []uint                         `json:"certificate_ids"`
	Group          string                         `json:"group"`
	State          common_type.ComponentStageType `json:"state"`
	CreatedAt      string                         `json:"created_at"`
	DeletedAt      string                         `json:"deleted_at"`
}

func BuildComponent(item models.Component) Component {
	certificates := ""
	certificateIDs := []uint{}
	if len(item.Certificates) > 0 {
		for _, one := range item.Certificates {
			certificates += one.Title + " "
			certificateIDs = append(certificateIDs, one.ID)
		}
		// 去掉末尾的空格
		certificates = certificates[:len(certificates)-1]
	}

	return Component{
		ID:             item.Model.ID,
		Name:           item.Name,
		PN:             item.PN,
		ManualNum:      item.Manual.Num,
		ModifyCount:    len(item.Modifies),
		Certificates:   certificates,
		CertificateIDs: certificateIDs,
		Group:          item.Group.Name,
		State:          item.State,
		CreatedAt:      utils.TimeFormat_YMD(item.CreatedAt),
		DeletedAt:      utils.TimeFormat_YMD(item.CreatedAt),
	}
}

func BuildComponents(items []models.Component) (components []Component) {
	for _, item := range items {
		component := BuildComponent(item)
		components = append(components, component)
	}
	return components
}
