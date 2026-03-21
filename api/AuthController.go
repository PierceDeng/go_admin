package api

import (
	resp "go_admin/model"
	reqVO "go_admin/model/reqVO"
	user "go_admin/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {

	var userLoginVO reqVO.UserLoginReqVO
	// 使用 ShouldBindJSON 绑定 JSON 并校验
	if err := c.ShouldBindJSON(&userLoginVO); err != nil {
		// 处理错误，如返回 400
		c.JSON(http.StatusOK, resp.Response{
			Code: 10001,
			Msg:  "json格式错误",
			Data: nil,
		})
		return
	}
	resp.Success(c, user.Login(userLoginVO))

}
