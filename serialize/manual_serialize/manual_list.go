package manual_serialize

import (
	"wild_goose_gin/models"
)

type Manual struct {
	ManualID uint   `json:"manual_id"`
	Num      string `json:"manual_num"`
	Version  string `json:"manual_version"`
}

func BuildManual(item models.Manual) Manual {
	return Manual{
		ManualID: item.ID,
		Num:      item.Num,
		Version:  item.Version,
	}
}
func BuildManuals(items []models.Manual) (manuals []Manual) {
	for _, item := range items {
		manual := BuildManual(item)
		manuals = append(manuals, manual)
	}
	return manuals
}
