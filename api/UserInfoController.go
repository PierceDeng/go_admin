package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/model/reqVO/user"
	userSerivce "go_admin/service/user"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	resp.Ok(c, userSerivce.GetUserInfo(userId.(uint64)))
}

func GetDeptTree(c *gin.Context) {
	dept := common.BindQuery[entity.SysDept](c)
	resp.Ok(c, userSerivce.GetDeptTree(dept))
}

func GetUserList(c *gin.Context) {
	userReqVO := common.BindQuery[user.SysUserReqVO](c)
	resp.OkWithWrapper(c, userSerivce.GetUserList(userReqVO))
}
