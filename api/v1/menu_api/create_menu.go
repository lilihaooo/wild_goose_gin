package menu_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
)

type CreateMenuRequest struct {
	ID       uint                 `json:"id" validate:"required" label:"ID"`
	Icon     string               `json:"icon"  label:"图标"`
	Path     string               `json:"index" validate:"required" label:"跳转路径"`
	Title    string               `json:"title" validate:"required" label:"标题"`
	Routes   []*int               `json:"routes" validate:"required" label:"路由"`
	ParentID uint                 `json:"parent_id"`
	Type     common_type.MenuType `json:"type"`
	Subs     []*models.Menu       `json:"subs"`
}

func (MenuApi) CreateMenu(c *gin.Context) {
	// todo 用于初始化菜单, 通过postApi将json菜单保存到数据库
	var req []CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "接收参数失败"+err.Error()) // todo err
		return
	}

	// 遍历 MenuItem 切片，保存到数据库
	for _, menuItem := range req {
		if err := global.DB.Model(models.Menu{}).Create(&menuItem).Error; err != nil {
			response.FailWithMsg(c, response.FAIL_OPER, err.Error())
			return
		}
	}

	response.OkWithMsg(c, "添加成功")
	//response.OkWithData(c, req)
}
