package auth

import (
	"go_admin/middleware/cache"
	"go_admin/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{Code: 9998, Msg: "Token未传递"})
			return
		}
		token = strings.ReplaceAll(token, "Bearer ", "")
		userId := cache.GetSysToken(token)
		if userId == 0 {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{Code: 9997, Msg: "请重新登录"})
			return
		}
		c.Set("userId", userId)
		c.Next()
	}

}
