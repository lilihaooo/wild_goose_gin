package material_serialize

import (
	"wild_goose_gin/models"
)

type ManualMaterial struct {
	ID            uint   `json:"id"`
	ImageNum      string `json:"image_num"`
	IsMustReplace string `json:"is_must_replace"`
	MaterialName  string `json:"material_name"` //批次号
}

func BuildMaterial(item models.ManualMaterial) ManualMaterial {
	isMustReplace := ""
	if item.IsMustReplace {
		isMustReplace = "是"
	}

	return ManualMaterial{
		ID:            item.ID,
		ImageNum:      item.ImageNum,
		IsMustReplace: isMustReplace,
		MaterialName:  item.Material.Name,
	}
}
func BuildManualMaterials(items []models.ManualMaterial) (manualMaterials []ManualMaterial) {
	for _, item := range items {
		manualMaterial := BuildMaterial(item)
		manualMaterials = append(manualMaterials, manualMaterial)
	}
	return manualMaterials
}
