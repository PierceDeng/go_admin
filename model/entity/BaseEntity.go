package entity

import "time"

type BaseEntity struct {
	CreateBy   string     `json:"createBy"`
	CreateTime *time.Time `json:"createTime"`
	UpdateBy   string     `json:"updateBy"`
	UpdateTime *time.Time `json:"updateTime"`
	Remark     string     `json:"remark"`
}
