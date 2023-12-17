package serializer

import (
	"wild_goose_gin/models"
	"wild_goose_gin/utils"
)

type Task struct {
	ID            uint   `json:"id"`
	TaskNum       string `json:"task_num"`
	CustomName    string `json:"custom_name"`
	ComponentName string `json:"component_name"`
	CreatedAt     string `json:"created_at"`
	PN            string `json:"pn"`
	SN            string `json:"sn"`
}

func BuildTask(item models.Task) Task {
	return Task{
		ID:            item.Model.ID,
		TaskNum:       item.TaskNum,
		CustomName:    item.Custom.Name,
		ComponentName: item.Component.Name,
		PN:            item.Component.PN,
		SN:            item.SN,
		CreatedAt:     utils.TimeFormat(item.CreatedAt),
	}
}

func BuildTasks(items []models.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
