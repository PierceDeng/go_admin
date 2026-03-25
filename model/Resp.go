package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	})
}

func OkWithWrapper(c *gin.Context, wrapper any) {
	c.JSON(http.StatusOK, wrapper)
}

func Fail(c *gin.Context, code int32, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	})
}
