package menu_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils"
)

type UpdateMenuRequest struct {
	ID       *uint                `json:"id" validate:"required" label:"ID"`
	ParentID uint                 `json:"parent_id"`
	Icon     string               `json:"icon"`
	Path     string               `json:"path"`
	Title    string               `json:"title" validate:"required" label:"标题"`
	RouteIDs []uint               `json:"route_ids"`
	Type     common_type.MenuType `json:"type"`
	Subs     []*models.Menu       `json:"subs"`
}

func (MenuApi) UpdateMenu(c *gin.Context) {
	var req UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "接收参数失败"+err.Error()) // todo err
		return
	}
	if vErr := utils.ZhValidate(&req); vErr != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, vErr)
		return
	}
	var menuModel models.Menu
	// 判断该记录是否存在
	menuModel.ID = req.ID
	menu, err := menuModel.GetOneRecordById()
	if err != nil {
		response.FailWithMsg(c, response.ERROR_NOT_EXIST_RECODE, "")
		return
	}
	menu.Title = req.Title
	// 0. 如果是root可以修改icon
	if req.ParentID == 0 {
		menu.Icon = req.Icon
	}
	// 2. 如果修改的是菜单: path, routes  // todo 菜单权限表移除
	//if req.Type == common_type.IsMenu {
	//	var routeModel models.Permission
	//	routes, err := routeModel.GetRouteListByIDs(req.RouteIDs)
	//	if err != nil {
	//		response.FailWithMsg(c, response.FAIL_OPER, "")
	//		return
	//	}
	//	menu.Path = req.Path
	//	//menu.Routes = routes
	//}
	// 修改菜单
	err = menu.SaveMenu()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	response.OkWithMsg(c, "修改成功")
}
