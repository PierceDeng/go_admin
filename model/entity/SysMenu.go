package entity

// SysMenu 菜单权限表（对应 Java 中的 SysMenu）
type SysMenu struct {
	BaseEntity // 嵌入基类字段

	MenuId     int64      `json:"menuId"`                   // 菜单ID
	MenuName   string     `json:"menuName" form:"menuName"` // 菜单名称（非空，长度≤50）
	ParentName string     `json:"parentName"`               // 父菜单名称
	ParentId   int64      `json:"parentId"`                 // 父菜单ID
	OrderNum   int        `json:"orderNum"`                 // 显示顺序（非空）
	Path       string     `json:"path"`                     // 路由地址（长度≤200）
	Component  string     `json:"component"`                // 组件路径（长度≤255）
	Query      string     `json:"query"`                    // 路由参数
	RouteName  string     `json:"routeName"`                // 路由名称（默认驼峰格式）
	IsFrame    string     `json:"isFrame"`                  // 是否为外链（0是 1否）
	IsCache    string     `json:"isCache"`                  // 是否缓存（0缓存 1不缓存）
	MenuType   string     `json:"menuType"`                 // 类型（M目录 C菜单 F按钮）（非空）
	Visible    string     `json:"visible"`                  // 显示状态（0显示 1隐藏）
	Status     string     `json:"status" form:"status"`     // 菜单状态（0正常 1停用）
	Perms      string     `json:"perms"`                    // 权限字符串（长度≤100）
	Icon       string     `json:"icon"`                     // 菜单图标
	Children   []*SysMenu `gorm:"-" json:"children"`        // 子菜单（递归嵌套）
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
