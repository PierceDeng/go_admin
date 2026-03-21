package exception

import (
	"fmt"
	resp "go_admin/model"

	"github.com/gin-gonic/gin"
)

func ExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if bizError, ok := err.(*BizError); ok {
					resp.Fail(c, bizError.Code, bizError.Msg)
				} else if paramError, ok := err.(*ParamError); ok {
					resp.Fail(c, paramError.Code, paramError.Msg)
				} else {
					fmt.Println("未知异常信息", err)
					resp.Fail(c, 9999, "系统错误")
				}
				c.Abort()
			}
		}()
		c.Next()
	}

}
