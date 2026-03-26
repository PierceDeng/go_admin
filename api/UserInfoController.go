package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/model/reqVO/user"
	userSerivce "go_admin/service/user"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserService *userSerivce.UserService
}

var UserController = userController{
	UserService: userSerivce.NewUserService(),
}

func (u userController) GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	resp.Ok(c, u.UserService.GetUserInfo(userId.(uint64)))
}

func (u userController) GetDeptTree(c *gin.Context) {
	dept, err := common.BindQuery[entity.SysDept](c)
	if err != nil {
		return
	}
	resp.Ok(c, u.UserService.GetDeptTree(dept))
}

func (u userController) GetUserList(c *gin.Context) {
	userReqVO, err := common.BindQuery[user.SysUserReqVO](c)
	if err != nil {
		return
	}
	resp.OkWithWrapper(c, u.UserService.GetUserList(userReqVO))
}
