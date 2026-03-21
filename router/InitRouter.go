package router

import (
	auth "go_admin/api"

	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {

	r.POST("/login", auth.UserLogin)
}
