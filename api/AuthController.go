package api

import (
	"go_admin/middleware/common"
	resp "go_admin/model"
	reqVO "go_admin/model/reqVO"
	user "go_admin/service/user"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	userLoginVO := common.BindJSON[reqVO.UserLoginReqVO](c)
	resp.Ok(c, user.Login(userLoginVO))
}
