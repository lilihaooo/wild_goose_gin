package menu_ser

import (
	"wild_goose_gin/models"
)

func (m MenuService) GetAllMenuTree(menus []*models.Menu) []*models.Menu {
	return buildMenuTree(menus, 0)
}

func buildMenuTree(allMenus []*models.Menu, parentID uint) []*models.Menu {
	var tree []*models.Menu

	for _, menu := range allMenus {
		if menu.ParentID == parentID {
			subMenu := buildMenuTree(allMenus, menu.ID)
			menu.Subs = subMenu
			tree = append(tree, menu)
		}
	}
	return tree
}
