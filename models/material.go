package models

import (
	"time"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/request"
)

type Material struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	PN           string    `json:"pn"`
	BN           string    `json:"bn"`          //批次号
	StorageNum   string    `json:"storage_num"` // 仓位号
	Count        uint      `json:"count"`
	Unit         string    `json:"unit"`         // 单位
	Manufacturer string    `json:"manufacturer"` // 制造商
	Price        uint      `json:"price"`        // 单价
	MinCount     uint      `json:"min_count"`    // 最少数量
	CreatedAt    time.Time `json:"created_at"`
	ExpiryAt     time.Time `json:"expiry_at"` // 过期时间
	Remark       string    `json:"remark"`    // 备注
}

func (m *Material) GetAllRecord(req *request.PaginationReq) (list []Material, count int64, err error) {
	query := global.DB.Model(m)
	if req.Keyword != "" {
		query.Where("pn like ?", "%"+req.Keyword+"%").Or("bn = ?", req.Keyword)
	}
	query.Count(&count)
	err = query.Offset(req.GetOffset()).Limit(req.PageSize).Find(&list).Error
	return list, count, err
}
