package entity

type SysUserPost struct {
	UserId uint64 `gorm:"user_id"`
	PostId int64  `gorm:"post_id"`
}

func (SysUserPost) TableName() string {
	return "sys_user_post"
}
