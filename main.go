package main

import (
	"go_admin/middleware/auth"
	"go_admin/middleware/exception"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
import "go_admin/router"
import "go_admin/config"

func init() {
	viper.SetConfigName("config")   // 文件名（不含扩展名）
	viper.SetConfigType("yaml")     // 文件格式
	viper.AddConfigPath("./config") // 查找路径
	viper.AddConfigPath(".")        // 可选，根目录兜底
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
}

func main() {

	config.InitDB()
	config.InitRedis()

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
