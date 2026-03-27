package user

import "go_admin/model/entity"

type QueryUserInfoRespVO struct {
	entity.SysUser
	Roles   []*entity.SysRole `gorm:"-" json:"roles"`
	Posts   []*entity.SysPost `gorm:"-" json:"posts"`
	PostIds []int64           `gorm:"-" json:"postIds"`
	RoleIds []int64           `gorm:"-" json:"roleIds"`
}
