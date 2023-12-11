package v1

import (
	"wild_goose_gin/api/v1/component_api"
	"wild_goose_gin/api/v1/custom_api"
	"wild_goose_gin/api/v1/manual_api"
	"wild_goose_gin/api/v1/menu_api"
	"wild_goose_gin/api/v1/modify_api"
	"wild_goose_gin/api/v1/route_api"
	"wild_goose_gin/api/v1/settings_api"
	"wild_goose_gin/api/v1/task_api"
	"wild_goose_gin/api/v1/user_api"
)

type ApiGroup struct {
	SettingsApi  settings_api.SettingsApi
	ComponentApi component_api.ComponentApi
	UserApi      user_api.UserApi
	TaskApi      task_api.TaskApi
	ManualApi    manual_api.ManualApi
	ModifyApi    modify_api.ModifyApi
	CustomApi    custom_api.CustomApi
	MenuApi      menu_api.MenuApi
	RouteApi     route_api.RouteApi
}

var ApiGroupApp = new(ApiGroup)
