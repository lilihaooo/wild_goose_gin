package models

import (
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/request"
)

type ManualMaterial struct {
	ID            uint `json:"id"`
	ImageNum      string
	IsMustReplace bool
	ManualID      uint `gorm:"foreignKey:ManualID;references:ID"`
	Manual        Manual
	MaterialID    uint `gorm:"foreignKey:MaterialID;references:ID"`
	Material      Material
}

func (m *ManualMaterial) GetListByManualID(pReq *request.PaginationReq) (list []ManualMaterial, count int64, err error) {
	limit := pReq.PageSize
	offset := pReq.GetOffset()
	db := global.DB.Model(m).Where("manual_id = ?", m.ManualID)
	err = db.Count(&count).Error
	err = db.Preload("Material").Limit(limit).Offset(offset).Find(&list).Error
	return
}
