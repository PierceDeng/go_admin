package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/service/menu"
	"strconv"

	"github.com/gin-gonic/gin"
)

type menuController struct {
	*menu.MenuService
}

var MenuController = menuController{
	MenuService: menu.NewMenuService(),
}

func (tis menuController) GetMenuList(c *gin.Context) {

	userId, _ := c.Get("userId")
	menu, err := common.BindQuery[entity.SysMenu](c)
	if err != nil {
		return
	}
	menuList := tis.MenuService.SelectList(menu, userId.(uint64))
	resp.Ok(c, menuList)
}

func (tis menuController) MenuInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resp.Ok(c, tis.MenuService.MenuInfo(id))
}

func (tis menuController) MenuTreeSelect(c *gin.Context) {
	userId, _ := c.Get("userId")
	menu, err := common.BindJSON[entity.SysMenu](c)
	if err != nil {
		return
	}
	menuList := tis.MenuService.SelectList(menu, userId.(uint64))
	resp.Ok(c, tis.MenuService.BuildMenuTree(menuList))
}

func (tis menuController) MenuDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resp.Ok(c, tis.MenuService.MenuDel(id))
}
