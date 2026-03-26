package dict

import (
	"go_admin/config"
	"go_admin/model/entity"
)

type DictService struct {
}

func NewDictService() *DictService {
	return &DictService{}
}

func (DictService) GetDictDateByType(dictTypeStr string) []*entity.SysDictData {

	var dictType []*entity.SysDictData
	config.DB.Where("status = '0' and dict_type = ?", dictTypeStr).Order("dict_sort asc").Find(&dictType)
	return dictType
}
