package user

type ResetUserPwdReqVO struct {
	UserId   uint64 `json:"userId"`
	Password string `json:"password"`
}
