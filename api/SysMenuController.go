package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	menuService "go_admin/service/menu"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMenuList(c *gin.Context) {

	userId, _ := c.Get("userId")
	menu := common.BindQuery[entity.SysMenu](c)
	menuList := menuService.SelectList(menu, userId.(uint64))
	resp.Ok(c, menuList)
}

func MenuInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resp.Ok(c, menuService.MenuInfo(id))
}

func MenuTreeSelect(c *gin.Context) {
	userId, _ := c.Get("userId")
	menu := common.BindJSON[entity.SysMenu](c)
	menuList := menuService.SelectList(menu, userId.(uint64))
	resp.Ok(c, menuService.BuildMenuTree(menuList))
}

func MenuDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resp.Ok(c, menuService.MenuDel(id))
}
