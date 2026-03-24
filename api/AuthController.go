package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/model/reqVO"
	menuService "go_admin/service/menu"
	"go_admin/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	userLoginVO := common.BindJSON[reqVO.UserLoginReqVO](c)
	resp.Ok(c, user.Login(userLoginVO))
}

func Logout(c *gin.Context) {
	resp.Ok(c, "")
}

func GetRouters(c *gin.Context) {

	userId, _ := c.Get("userId")
	menus := menuService.SelectMenuTreeByUserId(userId.(uint64))
	routerVOs := menuService.BuildMenus(menus)
	resp.Ok(c, routerVOs)

}

func GetMenuList(c *gin.Context) {

	userId, _ := c.Get("userId")
	menu := common.BindJSON[entity.SysMenu](c)
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
