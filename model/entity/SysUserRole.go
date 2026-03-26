package entity

type SysUserRole struct {
	UserId int `gorm:"user_id"`
	RoleId int `gorm:"role_id"`
}
