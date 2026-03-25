package entity

// SysDictData 字典数据表（对应 sys_dict_data）
type SysDictData struct {
	BaseEntity // 嵌入基类

	DictCode  int64  `json:"dictCode"`  // 字典编码
	DictSort  int64  `json:"dictSort"`  // 字典排序
	DictLabel string `json:"dictLabel"` // 字典标签（非空，长度≤100）
	DictValue string `json:"dictValue"` // 字典键值（非空，长度≤100）
	DictType  string `json:"dictType"`  // 字典类型（非空，长度≤100）
	CssClass  string `json:"cssClass"`  // 样式属性（其他样式扩展）
	ListClass string `json:"listClass"` // 表格字典样式
	IsDefault string `json:"isDefault"` // 是否默认（Y是 N否）
	Status    string `json:"status"`    // 状态（0正常 1停用）
}

// IsDefault 判断是否默认（对应 Java 中的 getDefault() 方法）
func (d *SysDictData) HasDefault() bool {
	return d.IsDefault == "Y"
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
