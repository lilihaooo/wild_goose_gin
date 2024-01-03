package component_ser

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

func (s ComponentService) BatchDeleteComponentAndAssociation(ids []int64) error {
	tx := global.DB.Begin()
	// 清除关联数据1
	if err := tx.Exec("DELETE FROM component_certificate WHERE component_id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 清除关联数据2
	if err := tx.Exec("DELETE FROM modify WHERE component_id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	var componentModel models.Component
	if err := tx.Where("id IN ?", ids).Delete(&componentModel).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
