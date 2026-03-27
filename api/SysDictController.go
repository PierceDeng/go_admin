package api

import (
	resp "go_admin/model"
	"go_admin/service/dict"

	"github.com/gin-gonic/gin"
)

type sysDictController struct {
	*dict.DictService
}

var SysDictController = &sysDictController{
	dict.NewDictService(),
}

func (tis *sysDictController) GetDictDateInfo(c *gin.Context) {

	typeId := c.Param("type")
	resp.Ok(c, tis.DictService.GetDictDateByType(typeId))
}
