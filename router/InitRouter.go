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

	authGroup.GET("/getInfo", api.UserController.GetUserInfo)
	authGroup.GET("/system/user/deptTree", api.UserController.GetDeptTree)
	authGroup.GET("/system/user/list", api.UserController.GetUserList)
	authGroup.PUT("/system/user/changeStatus", api.UserController.ChangeUserStatus)
	authGroup.GET("/system/user/:userId", api.UserController.QueryUser)
	authGroup.GET("/getRouters", api.AuthController.GetRouters)
	authGroup.GET("/system/menu/list", api.MenuController.GetMenuList)
	authGroup.GET("/system/menu/:id", api.MenuController.MenuInfo)
	authGroup.DELETE("/system/menu/:id", api.MenuController.MenuDel)
	authGroup.GET("/system/menu/tree", api.MenuController.MenuTreeSelect)
	authGroup.GET("/system/dict/data/type/:type", api.SysDictController.GetDictDateInfo)

}
