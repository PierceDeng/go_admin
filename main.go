package main

import (
	"go_admin/middleware/exception"

	"github.com/gin-gonic/gin"
)
import "go_admin/router"
import "go_admin/config"

func main() {

	config.InitDB()
	defer config.CloseDB()

	config.InitRedis()
	defer config.CloseRedis()

	r := gin.Default()

	r.Use(exception.ExceptionHandler())

	router.LoadRouter(r)

	r.Run(":9999")
}
