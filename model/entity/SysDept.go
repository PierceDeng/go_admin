package entity

type SysDept struct {
	BaseEntity // 嵌入基类

	DeptId     int64      `json:"deptId"`            // 部门ID
	ParentId   int64      `json:"parentId"`          // 父部门ID
	Ancestors  string     `json:"ancestors"`         // 祖级列表
	DeptName   string     `json:"deptName"`          // 部门名称（非空，长度≤30）
	OrderNum   int        `json:"orderNum"`          // 显示顺序（非空）
	Leader     string     `json:"leader"`            // 负责人
	Phone      string     `json:"phone"`             // 联系电话（长度≤11）
	Email      string     `json:"email"`             // 邮箱（格式校验，长度≤50）
	Status     string     `json:"status"`            // 部门状态（0正常 1停用）
	DelFlag    string     `json:"delFlag"`           // 删除标志（0存在 2删除）
	ParentName string     `json:"parentName"`        // 父部门名称（仅用于显示）
	Children   []*SysDept `gorm:"-" json:"children"` // 子部门列表
}

func (SysDept) TableName() string {
	return "sys_dept"
}
