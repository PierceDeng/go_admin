package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/reqVO"
	"go_admin/service/menu"
	"go_admin/service/user"

	"github.com/gin-gonic/gin"
)

type authController struct {
	*user.UserService
	*menu.MenuService
}

var AuthController = &authController{
	UserService: user.NewUserService(),
	MenuService: menu.NewMenuService(),
}

func (a *authController) UserLogin(c *gin.Context) {
	userLoginVO, err := common.BindJSON[reqVO.UserLoginReqVO](c)
	if err != nil {
		return
	}
	resp.Ok(c, a.UserService.Login(userLoginVO))
}

func (a *authController) Logout(c *gin.Context) {
	resp.Ok(c, "")
}

func (a *authController) GetRouters(c *gin.Context) {

	userId, _ := c.Get("userId")
	menus := a.MenuService.SelectMenuTreeByUserId(userId.(uint64))
	routerVOs := a.MenuService.BuildMenus(menus)
	resp.Ok(c, routerVOs)

}
