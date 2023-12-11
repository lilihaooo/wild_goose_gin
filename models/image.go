package models

import (
	"gorm.io/gorm"
	"os"
	"time"
	"wild_goose_gin/global"
	"wild_goose_gin/models/common_type"
)

type Image struct {
	ID        uint                  `json:"id"`
	Path      string                `gorm:"comment:图片路径" json:"path"`                         // 图片路径
	Hash      string                `gorm:"comment:图片的hash值，用于判断重复图片" json:"hash"`            // 图片的hash值，用于判断重复图片
	Name      string                `gorm:"size:38;comment:图片名称" json:"name"`                 // 图片名称
	Type      common_type.ImageType `gorm:"comment:图片类型 1: 本地, 2: 服务器;default:1" json:"type"` // 图片类型 1: 本地, 2: 服务器
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func (b *Image) BeforeDelete(tx *gorm.DB) (err error) {
	if b.Type == common_type.Local {
		if err = os.Remove(b.Path); err != nil {
			global.Logrus.Error(err)
			return err
		}
	}
	return nil
}

func (b *Image) GetImageIDByHash() uint {
	if err := global.DB.Where("hash = ?", b.Hash).Find(b).Error; err != nil {
		return 0
	}
	return b.ID
}
