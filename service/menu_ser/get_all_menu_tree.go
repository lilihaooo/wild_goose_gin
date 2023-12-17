package menu_ser

import (
	"wild_goose_gin/models"
)

type Tree struct {
	ID       uint    `json:"id"`
	Icon     string  `json:"icon"`
	Path     string  `json:"path"`
	Title    string  `json:"title"`
	ParentID uint    `json:"parent_id,omitempty"`
	Type     uint    `json:"type"`
	RouteIDs []uint  `json:"route_ids"`
	Subs     []*Tree `json:"subs"`
}

func (m MenuService) GetAllMenuTree(menus []*models.Menu) []*models.Menu {
	return buildMenuTree(menus, 0)
}

func buildMenuTree(allMenus []*models.Menu, parentID uint) []*models.Menu {
	var tree []*models.Menu

	for _, menu := range allMenus {
		if menu.ParentID == nil {
			value := uint(0)
			menu.ParentID = &value
		}
		if *menu.ParentID == parentID {
			subMenu := buildMenuTree(allMenus, *menu.ID)
			menu.Subs = subMenu
			tree = append(tree, menu)
		}
	}
	return tree
}


