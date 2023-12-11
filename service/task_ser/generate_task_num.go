package task_ser

import (
	"fmt"
	"strconv"
	"time"
	"wild_goose_gin/models"
)

func (s TaskService) GenerateTaskNum(group models.Group) (string, error) {
	var task models.Task
	thisMonthCount, err := task.GetGroupThisMonthCount(group.ID)
	if err != nil {
		return "", err
	}
	waterNum := fmt.Sprintf("%04d", thisMonthCount+1)

	// 获取当前时间
	now := time.Now()
	// 获取当前年份的后两位
	yearLastTwoDigits := strconv.Itoa(now.Year() % 100)
	// 获取当前月份，并使用0填充单数月份
	month := now.Month()
	monthWithZeroPadding := fmt.Sprintf("%02d", month)
	taskNum := group.Prefix + yearLastTwoDigits + monthWithZeroPadding + waterNum
	return taskNum, err
}
