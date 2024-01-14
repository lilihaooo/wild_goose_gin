package material_serialize

import (
	"wild_goose_gin/models"
	"wild_goose_gin/utils"
)

type Material struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	PN           string `json:"pn"`
	BN           string `json:"bn"` //批次号
	StorageNum   string `json:"storage_num"`
	Count        uint   `json:"count"`
	Unit         string `json:"unit"`
	Manufacturer string `json:"manufacturer"`
	Price        uint   `json:"price"`
	MinCount     uint   `json:"min_count"`
	CreatedAt    string `json:"created_at"`
	ExpiryAt     string `json:"expiry_at"`
}

func BuildMaterial(item models.Material) Material {

	return Material{
		ID:           item.ID,
		Name:         item.Name,
		PN:           item.PN,
		BN:           item.BN,
		StorageNum:   item.StorageNum,
		Count:        item.Count,
		Unit:         item.Unit,
		Manufacturer: item.Manufacturer,
		Price:        item.Price,
		MinCount:     item.MinCount,
		CreatedAt:    utils.TimeFormat_YMD(item.CreatedAt),
		ExpiryAt:     utils.TimeFormat_YMD(item.ExpiryAt),
	}
}
func BuildMaterials(items []models.Material) (materials []Material) {
	for _, item := range items {
		material := BuildMaterial(item)
		materials = append(materials, material)
	}
	return materials
}
