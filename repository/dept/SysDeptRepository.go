package dept

import (
	"go_admin/config"
	"go_admin/model/entity"
)

func SelectDeptList(sysDept *entity.SysDept) (respList []*entity.SysDept) {

	config.DB.Where(sysDept).Where("del_flag = 0").Find(&respList)
	return respList
}
