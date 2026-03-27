package entity

type SysUserRole struct {
	UserId uint64 `gorm:"user_id"`
	RoleId int64  `gorm:"role_id"`
}

func (r SysUserRole) TableName() string { return "sys_user_role" }
