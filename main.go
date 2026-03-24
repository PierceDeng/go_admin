package main

import (
	"go_admin/middleware/auth"
	"go_admin/middleware/exception"

	"github.com/gin-gonic/gin"
)
import "go_admin/router"
import "go_admin/config"

func main() {

	config.InitDB()
	config.InitRedis()

	defer config.CloseDB()
	defer config.CloseRedis()

	r := gin.Default()
	r.Use(exception.ExceptionHandler())

	authGroup := r.Group("/", auth.AuthFilter())
	noAuthGroup := r.Group("/")

	router.LoadRouter(noAuthGroup)
	router.LoadAuthRouter(authGroup)

	err := r.Run(":9999")
	if err != nil {
		return
	}
}
