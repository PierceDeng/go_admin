package entity

type SysRole struct {
	BaseEntity // 嵌入基类字段

	RoleId            int64    `json:"roleId"`               // 角色ID
	RoleName          string   `json:"roleName"`             // 角色名称（长度≤30）
	RoleKey           string   `json:"roleKey"`              // 角色权限字符串（长度≤100）
	RoleSort          int      `json:"roleSort"`             // 显示顺序
	DataScope         string   `json:"dataScope"`            // 数据范围（1-5）
	MenuCheckStrictly bool     `json:"menuCheckStrictly"`    // 菜单树选择项是否关联显示
	DeptCheckStrictly bool     `json:"deptCheckStrictly"`    // 部门树选择项是否关联显示
	Status            string   `json:"status"`               // 状态（0正常 1停用）
	DelFlag           string   `json:"delFlag"`              // 删除标志（0存在 2删除）
	Flag              bool     `json:"flag"`                 // 用户是否存在此角色标识，默认false
	MenuIds           []int64  `gorm:"-" json:"menuIds"`     // 菜单组
	DeptIds           []int64  `gorm:"-" json:"deptIds"`     // 部门组（数据权限）
	Permissions       []string `gorm:"-" json:"permissions"` // 角色菜单权限
}

func (r SysRole) IsAdmin() bool {
	return 1 == r.RoleId
}

func (r SysRole) TableName() string {
	return "sys_role"
}
