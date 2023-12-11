package service

import (
	"wild_goose_gin/service/common/redis_ser"
	"wild_goose_gin/service/component_ser"
	"wild_goose_gin/service/image_ser"
	"wild_goose_gin/service/menu_ser"
	"wild_goose_gin/service/task_ser"
)

type Service struct {
	image_ser.ImageService
	redis_ser.RedisService
	component_ser.ComponentService
	task_ser.TaskService
	menu_ser.MenuService
}

var AppService = new(Service)
