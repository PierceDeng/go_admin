package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/reqVO"
	menuService "go_admin/service/menu"
	"go_admin/service/user"

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
