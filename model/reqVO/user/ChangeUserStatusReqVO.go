package user

type ChangeUserStatusReqVo struct {
	UserId uint64 `json:"userId"`
	Status string `json:"status"`
}
