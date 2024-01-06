package serror

type SError struct {
	IsSystem bool
	Code     string
	Msg      string
}

func (e *SError) Error() string {
	return e.Msg
}

func NewSError(code string, msg string) *SError {
	return &SError{
		IsSystem: false,
		Code:     code,
		Msg:      msg,
	}
}

func NewSystemSError(msg string) *SError {
	return &SError{
		IsSystem: true,
		Code:     ErrSystemInternal,
		Msg:      msg,
	}
}
