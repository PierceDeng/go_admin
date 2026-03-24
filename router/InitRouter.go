package router

import (
	api "go_admin/api"

	"github.com/gin-gonic/gin"
)

func LoadRouter(noAuthGroup *gin.RouterGroup) {

	noAuthGroup.POST("/login", api.UserLogin)
	noAuthGroup.POST("/logout", api.Logout)
}

func LoadAuthRouter(authGroup *gin.RouterGroup) {

	authGroup.GET("/getInfo", api.GetUserInfo)
	authGroup.GET("/getRouters", api.GetRouters)
	authGroup.GET("/system/menu/list", api.GetMenuList)
}
