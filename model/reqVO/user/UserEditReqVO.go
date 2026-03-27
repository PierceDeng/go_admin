package user

import "go_admin/model/entity"

type UserEditReqVO struct {
	entity.SysUser
	RoleIds []int64 `gorm:"-" json:"roleIds"`
	PostIds []int64 `gorm:"-" json:"postIds"`
}
