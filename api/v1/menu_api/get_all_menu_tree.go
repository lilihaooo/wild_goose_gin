package menu_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
)

func (MenuApi) GetAllMenuTree(c *gin.Context) {
	var model models.Menu
	menus, err := model.GetAllRecord()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	menuService := service.AppService.MenuService
	tree := menuService.GetAllMenuTree(menus)
	response.OkWithData(c, tree)
}
