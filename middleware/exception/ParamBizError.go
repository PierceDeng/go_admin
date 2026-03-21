package exception

type ParamError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ParamError) error() string {
	return e.Msg
}

func NewParamException(code int32, msg string) *ParamError {

	return &ParamError{
		Code: code,
		Msg:  msg,
	}

}
