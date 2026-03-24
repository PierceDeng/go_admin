package RespVO

import "go_admin/model/entity"

type UserInfoRespVO struct {
	User               entity.SysUser `json:"user"`
	Roles              []string       `json:"roles"`
	Permissions        []string       `json:"permissions"`
	IsDefaultModifyPwd bool           `json:"isDefaultModifyPwd"`
	IsPasswordExpired  bool           `json:"isPasswordExpired"`
}
