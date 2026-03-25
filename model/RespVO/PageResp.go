package RespVO

// TableDataInfo 表格分页数据对象
type PageResp[T any] struct {
	Total int64  `json:"total"` // 总记录数
	Rows  []T    `json:"rows"`  // 列表数据
	Code  int    `json:"code"`  // 消息状态码
	Msg   string `json:"msg"`   // 消息内容
}

// NewTableDataInfo 创建分页数据对象
func NewTableDataInfo[T any](list []T, total int64) *PageResp[T] {
	return &PageResp[T]{
		Total: total,
		Rows:  list,
	}
}
