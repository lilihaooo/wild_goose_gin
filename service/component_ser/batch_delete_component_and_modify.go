package component_ser

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

func (s ComponentService) BatchDeleteComponentAndModify(ids []int64) error {
	db := global.DB
	tx := db.Begin()
	var componentModel models.Component
	if err := tx.Where("id IN ?", ids).Delete(&componentModel).Error; err != nil {
		tx.Rollback()
		return err
	}
	var modifyModel models.Modify
	for _, componentID := range ids {
		if err := tx.Where("component_id = ?", componentID).Delete(&modifyModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
