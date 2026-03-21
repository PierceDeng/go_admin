package exception

type BizError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BizError) error() string {
	return e.Msg
}

func NewBizException(code int32, msg string) *BizError {

	return &BizError{
		Code: code,
		Msg:  msg,
	}

}
