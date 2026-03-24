package menu

import (
	menuConst "go_admin/middleware/common"
	"go_admin/model/RespVO/menu"
	"go_admin/model/entity"
	mRepository "go_admin/repository/menu"
	"go_admin/utils"
	"strings"
)

const MENU_ROOT_ID = 0

func BuildMenus(menus []*entity.SysMenu) []*menu.RouterVO {

	var m []*menu.RouterVO
	for _, item := range menus {
		var router menu.RouterVO
		router.Hidden = "1" == item.Visible
		router.Name = getRouteName(*item)
		router.Path = getRouterPath(*item)
		router.Component = getComponent(*item)
		router.Query = item.Query
		router.Meta = &menu.MetaVo{
			Title:   item.MenuName,
			Icon:    item.Icon,
			NoCache: "1" == item.IsCache,
			Link:    item.Path,
		}
		children := item.Children
		if len(children) > 0 && menuConst.TYPE_DIR == item.MenuType {
			router.AlwaysShow = new(true)
			router.Redirect = "NoRedirect"
			router.Children = BuildMenus(children)
		} else if isMenuFrame(*item) {
			router.Meta = nil
			var childrenList []*menu.RouterVO
			var child menu.RouterVO
			child.Path = item.Path
			child.Component = item.Component
			child.Name = getRouteNameStr(item.RouteName, item.Path)
			child.Meta = &menu.MetaVo{
				Title:   item.MenuName,
				Icon:    item.Icon,
				NoCache: "1" == item.IsCache,
				Link:    item.Path,
			}
			child.Query = item.Query
			childrenList = append(childrenList, &child)
			router.Children = childrenList
		} else if item.ParentId == MENU_ROOT_ID && isInnerLink(*item) {

			router.Meta = &menu.MetaVo{
				Title: item.MenuName,
				Icon:  item.Icon,
			}
			router.Path = "/"

			var childrenList []*menu.RouterVO
			var child menu.RouterVO
			child.Path = item.Path
			child.Component = item.Component
			child.Name = getRouteNameStr(item.RouteName, item.Path)
			child.Meta = &menu.MetaVo{
				Title:   item.MenuName,
				Icon:    item.Icon,
				NoCache: "1" == item.IsCache,
				Link:    item.Path,
			}
			child.Query = item.Query
			childrenList = append(childrenList, &child)
			router.Children = childrenList
		}

		m = append(m, &router)
	}

	return m

}

func getComponent(sysMenu entity.SysMenu) string {

	var component = menuConst.LAYOUT
	if sysMenu.Component != "" && !isMenuFrame(sysMenu) {
		component = sysMenu.Component
	} else if sysMenu.Component == "" && sysMenu.ParentId != MENU_ROOT_ID && isInnerLink(sysMenu) {
		component = menuConst.INNER_LINK
	} else if sysMenu.Component == "" && isParentView(sysMenu) {
		component = menuConst.PARENT_VIEW
	}
	return component
}

func isParentView(sysMenu entity.SysMenu) bool {
	return sysMenu.ParentId != MENU_ROOT_ID && menuConst.TYPE_DIR == sysMenu.MenuType
}

func SelectMenuTreeByUserId(userId uint64) (menus []*entity.SysMenu) {

	if IsAdmin(userId) {
		menus = mRepository.SelectMenuTreeAll()
	} else {
		menus = mRepository.SelectMenuTreeByUserId(userId)
	}
	return getChildPerms(menus, MENU_ROOT_ID)
}

func IsAdmin(userId uint64) bool {
	return 1 == userId
}

func getChildPerms(menus []*entity.SysMenu, parentId int64) []*entity.SysMenu {

	var respList []*entity.SysMenu

	if menus != nil {
		for _, sysMenu := range menus {
			if sysMenu.ParentId == parentId {
				recursionFn(menus, sysMenu)
				respList = append(respList, sysMenu)
			}
		}
	}

	return respList
}

func recursionFn(menus []*entity.SysMenu, t *entity.SysMenu) {

	var children = getChildList(menus, t)
	t.Children = children

	for _, child := range children {
		if hasChild(menus, child) {
			recursionFn(menus, child)
		}
	}
}

func getChildList(menus []*entity.SysMenu, t *entity.SysMenu) []*entity.SysMenu {

	var tList []*entity.SysMenu
	for _, n := range menus {
		if n.ParentId == t.MenuId {
			tList = append(tList, n)
		}
	}
	return tList
}

func hasChild(menus []*entity.SysMenu, t *entity.SysMenu) bool {
	return len(getChildList(menus, t)) > 0
}

func getRouteName(sysMenu entity.SysMenu) string {
	if isMenuFrame(sysMenu) {
		return ""
	}
	return getRouteNameStr(sysMenu.RouteName, sysMenu.Path)
}

func getRouteNameStr(name string, path string) string {
	if name != "" {
		return name
	}
	return path
}

func isMenuFrame(sysMenu entity.SysMenu) bool {
	return sysMenu.ParentId == MENU_ROOT_ID && menuConst.TYPE_MENU == sysMenu.MenuType && sysMenu.IsFrame == menuConst.NO_FRAME

}

func getRouterPath(sysMenu entity.SysMenu) string {

	routerPath := sysMenu.Path
	if sysMenu.ParentId != MENU_ROOT_ID && isInnerLink(sysMenu) {
		routerPath = innerLinkReplaceEach(&routerPath)
		if MENU_ROOT_ID == sysMenu.ParentId && menuConst.TYPE_DIR == sysMenu.MenuType && menuConst.NO_FRAME == sysMenu.IsFrame {
			routerPath = "/" + routerPath
		} else if isMenuFrame(sysMenu) {
			routerPath = "/"
		}
	}
	return routerPath
}

func isInnerLink(sysMenu entity.SysMenu) bool {

	return sysMenu.IsFrame == menuConst.NO_FRAME && strings.Contains(sysMenu.Path, "http")
}

func innerLinkReplaceEach(path *string) string {
	return strings.ReplaceAll(*path, "http|https|www", "")
}

func SelectMenuPermsByRoleId(roleId int64) []string {

	menus := mRepository.SelectMenuPermsByRoleId(roleId)
	menus = utils.UniqueStrings(menus)

	var menusSet []string
	for _, str := range menus {
		if str != "" {
			menusSet = append(menusSet, strings.Split(str, ",")...)
		}
	}
	return menusSet
}

func SelectMenuPermsByUserId(id uint64) []string {
	perms := mRepository.SelectMenuPermsByUserId(id)
	perms = utils.UniqueStrings(perms)

	var permsSet []string
	for _, str := range perms {
		if str != "" {
			permsSet = append(permsSet, strings.Split(str, ",")...)
		}
	}
	return permsSet
}
