package api

import (
	resp "go_admin/model"
	userSerivce "go_admin/service/user"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	resp.Ok(c, userSerivce.GetUserInfo(userId.(uint64)))
}
