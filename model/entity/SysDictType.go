package entity

// SysDictType 字典类型表（对应 sys_dict_type）
type SysDictType struct {
	BaseEntity // 嵌入基类

	DictId   int64  `json:"dictId"`   // 字典主键
	DictName string `json:"dictName"` // 字典名称（非空，长度≤100）
	DictType string `json:"dictType"` // 字典类型（非空，长度≤100，正则：字母开头，小写字母+数字+下划线）
	Status   string `json:"status"`   // 状态（0正常 1停用）
}

// DictTypeRegexp 返回字典类型的正则表达式（用于验证）
func (SysDictType) DictTypeRegexp() string {
	return `^[a-z][a-z0-9_]*$`
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
