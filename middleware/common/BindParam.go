package common

import (
	resp "go_admin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON[T any](c *gin.Context) *T {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, resp.Response{Code: 10001, Msg: "参数格式错误", Data: gin.H{}})
		return nil
	}
	return &req
}

func BindQuery[T any](c *gin.Context) *T {
	var req T
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, resp.Response{Code: 10001, Msg: "参数格式错误", Data: gin.H{}})
		return nil
	}
	return &req
}
