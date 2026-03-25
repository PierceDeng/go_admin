package api

import (
	resp "go_admin/model"
	dictService "go_admin/service/dict"

	"github.com/gin-gonic/gin"
)

func GetDictDateInfo(c *gin.Context) {

	typeId := c.Param("type")
	resp.Ok(c, dictService.GetDictDateByType(typeId))
}
