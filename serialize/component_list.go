package serializer

import (
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"
)

type Component struct {
	ID            uint                           `json:"id"`
	Name          string                         `json:"name"`
	PN            string                         `json:"pn"`
	ManualNum     string                         `json:"manual_num"`
	ModifiesCount int                            `json:"modifies_count"`
	Group         string                         `json:"group"`
	State         common_type.ComponentStageType `json:"state"`
	CreatedAt     string                         `json:"created_at"`
	DeletedAt     string                         `json:"deleted_at"`
}

func BuildComponent(item models.Component) Component {
	return Component{
		ID:            item.Model.ID,
		Name:          item.Name,
		PN:            item.PN,
		ManualNum:     item.Manual.Num,
		ModifiesCount: len(item.Modifies),
		Group:         item.Group.Name,
		State:         item.State,
		CreatedAt:     utils.TimeFormat(item.CreatedAt),
		DeletedAt:     utils.TimeFormat(item.CreatedAt),
	}
}

func BuildComponents(items []models.Component) (components []Component) {
	for _, item := range items {
		component := BuildComponent(item)
		components = append(components, component)
	}
	return components
}
