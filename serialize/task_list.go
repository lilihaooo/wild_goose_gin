package serializer

import (
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"
)

type Task struct {
	ID            uint   `json:"id"`
	TaskNum       string `json:"task_num"`
	CustomName    string `json:"custom_name"`
	ComponentName string `json:"component_name"`

	Demand          common_type.TaskDemandType `json:"demand"`
	Certificates    string                     `json:"certificates"`
	ModifyCount     int                        `json:"modify_count"`
	Node            common_type.TaskNodeType   `json:"node"`
	Share           common_type.TaskShareType  `json:"share"`
	PlanReleaseDate string                     `json:"plan_release_date"`

	CreatedAt string `json:"created_at"`
	PN        string `json:"pn"`
	SN        string `json:"sn"`
}

func BuildTask(item models.Task) Task {
	certificates := ""
	if len(*item.Certificates) > 0 {
		for _, one := range *item.Certificates {
			certificates += one.Title + " "
		}
		// 去掉末尾的空格
		certificates = certificates[:len(certificates)-1]
	}

	return Task{
		ID:            item.Model.ID,
		TaskNum:       item.TaskNum,
		CustomName:    item.Custom.Name,
		ComponentName: item.Component.Name,
		PN:            item.Component.PN,
		SN:            item.SN,

		Demand:          item.Demand,
		Certificates:    certificates,
		ModifyCount:     len(item.Modifies),
		Node:            item.Node,
		Share:           item.Share,
		PlanReleaseDate: utils.TimeFormat_YMD(item.PlanReleaseDate),
		CreatedAt:       utils.TimeFormat_YMD(item.CreatedAt),
	}
}

func BuildTasks(items []models.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
