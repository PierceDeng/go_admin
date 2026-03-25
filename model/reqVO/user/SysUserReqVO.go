package user

import (
	"go_admin/model/entity"
	"go_admin/model/reqVO"
)

type SysUserReqVO struct {
	reqVO.PageEntity
	entity.SysUser

	Param struct {
		BeginTime string `json:"beginTime"`
		EndTime   string `json:"endTime"`
	}
}
