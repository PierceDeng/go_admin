package router

import (
	api "go_admin/api"

	"github.com/gin-gonic/gin"
)

func LoadRouter(noAuthGroup *gin.RouterGroup) {

	noAuthGroup.POST("/login", api.AuthController.UserLogin)
	noAuthGroup.POST("/logout", api.AuthController.Logout)
}

func LoadAuthRouter(authGroup *gin.RouterGroup) {

	authGroup.GET("/getInfo", api.UserControl.GetUserInfo)
	authGroup.GET("/system/user/deptTree", api.UserControl.GetDeptTree)
	authGroup.GET("/system/user/list", api.UserControl.GetUserList)
	authGroup.PUT("/system/user/changeStatus", api.UserControl.ChangeUserStatus)
	authGroup.GET("/system/user/:userId", api.UserControl.QueryUser)
	authGroup.GET("/system/user/", api.UserControl.QueryUser)
	authGroup.PUT("/system/user", api.UserControl.UpdateUser)
	authGroup.GET("/getRouters", api.AuthController.GetRouters)
	authGroup.GET("/system/menu/list", api.MenuController.GetMenuList)
	authGroup.GET("/system/menu/:id", api.MenuController.MenuInfo)
	authGroup.DELETE("/system/menu/:id", api.MenuController.MenuDel)
	authGroup.GET("/system/menu/tree", api.MenuController.MenuTreeSelect)
	authGroup.GET("/system/dict/data/type/:type", api.SysDictController.GetDictDateInfo)

}
