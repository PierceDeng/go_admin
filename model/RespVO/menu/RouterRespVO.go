package menu

// RouterVo 路由配置
type RouterVO struct {
	Name       string      `json:"name"`                 // 路由名字
	Path       string      `json:"path"`                 // 路由地址
	Hidden     bool        `json:"hidden"`               // 是否隐藏路由（侧边栏不显示）
	Redirect   string      `json:"redirect,omitempty"`   // 重定向地址
	Component  string      `json:"component"`            // 组件地址
	Query      string      `json:"query,omitempty"`      // 路由参数
	AlwaysShow *bool       `json:"alwaysShow,omitempty"` // 是否总是显示（当子路由>1时自动嵌套）
	Meta       *MetaVo     `json:"meta"`                 // 路由元信息
	Children   []*RouterVO `json:"children,omitempty"`   // 子路由
}

// MetaVo 路由元信息
type MetaVo struct {
	Title   string `json:"title"`          // 菜单标题
	Icon    string `json:"icon"`           // 菜单图标
	NoCache bool   `json:"noCache"`        // 是否缓存
	Link    string `json:"link,omitempty"` // 外链地址
}
