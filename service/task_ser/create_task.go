package task_ser

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

func (s TaskService) CreateTask(task *models.Task, incomeTotal, claimTotal int) error {
	var component models.Component
	component.ID = task.ComponentID
	// 开启事务同时修改component数据, 和添加任务
	tx := global.DB.Begin()
	err := tx.First(&component).Set("gorm:query_option", "FOR UPDATE").Error // 加锁
	if err != nil {
		tx.Rollback()
		return err
	}
	component.IncomeTotal = incomeTotal
	component.ClaimTotal = claimTotal
	err = tx.Save(&component).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(task).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
