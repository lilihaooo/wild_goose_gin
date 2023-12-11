package menu_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils/jwts"
)

func (MenuApi) GetRoleMenuTree(c *gin.Context) {
	var user models.User
	payload := c.MustGet("user").(*jwts.Payload)
	userID := payload.UserID
	user.ID = userID
	roleID, err := user.GetUserRoleID()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	var role models.Role
	role.ID = roleID
	roleMenus, err := role.GetRoleMenuList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	menuService := service.AppService.MenuService
	tree := menuService.GetAllMenuTree(roleMenus)
	response.OkWithData(c, tree)
}
