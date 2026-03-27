package entity

// SysPost 岗位表 sys_post
type SysPost struct {
	PostId   int64  `json:"postId" gorm:"primaryKey;autoIncrement;column:post_id;comment:岗位序号"`
	PostCode string `json:"postCode" gorm:"column:post_code;type:varchar(64);not null;comment:岗位编码" validate:"required,max=64"`
	PostName string `json:"postName" gorm:"column:post_name;type:varchar(50);not null;comment:岗位名称" validate:"required,max=50"`
	PostSort int32  `json:"postSort" gorm:"column:post_sort;type:int(4);not null;comment:岗位排序" validate:"required"`
	Status   string `json:"status" gorm:"column:status;type:char(1);default:0;comment:状态（0正常 1停用）"`
	Flag     bool   `json:"flag" gorm:"-"` // 忽略数据库映射

	BaseEntity
}

// TableName 指定表名
func (SysPost) TableName() string {
	return "sys_post"
}
