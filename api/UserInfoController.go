package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/model/reqVO/user"
	userSerivce "go_admin/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *userSerivce.UserService
}

var UserControl = NewUserController()

func NewUserController() *UserController {
	return &UserController{
		UserService: userSerivce.NewUserService(),
	}
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	resp.Ok(c, u.UserService.GetUserInfo(userId.(uint64)))
}

func (u *UserController) GetDeptTree(c *gin.Context) {
	dept, err := common.BindQuery[entity.SysDept](c)
	if err != nil {
		return
	}
	resp.Ok(c, u.UserService.GetDeptTree(dept))
}

func (u *UserController) GetUserList(c *gin.Context) {
	userReqVO, err := common.BindQuery[user.SysUserReqVO](c)
	if err != nil {
		return
	}
	resp.OkWithWrapper(c, u.UserService.GetUserList(userReqVO))
}

func (u *UserController) ChangeUserStatus(c *gin.Context) {
	reqVO, err := common.BindJSON[user.ChangeUserStatusReqVo](c)
	if err != nil {
		return
	}
	resp.Ok(c, u.UserService.ChangeUserStatus(reqVO))
}

func (u *UserController) QueryUser(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("userId"))
	resp.Ok(c, u.UserService.QueryUser(userId))
}

func (u *UserController) UpdateUser(c *gin.Context) {

	reqVO, _ := common.BindJSON[user.UserEditReqVO](c)
	resp.Ok(c, u.UserService.UpdateUser(c, reqVO))
}
