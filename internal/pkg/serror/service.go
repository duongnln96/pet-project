package serror

type SError struct {
	IsSystem bool
	IsAuth   bool
	Code     string
	Msg      string
}

func (e *SError) Error() string {
	return e.Msg
}

func NewSError(code string, msg string) *SError {
	return &SError{
		Code: code,
		Msg:  msg,
	}
}

func NewSystemSError(msg string) *SError {
	return &SError{
		IsSystem: true,
		Code:     ErrSystemInternal,
		Msg:      msg,
	}
}

func NewAuthError(msg string) *SError {
	return &SError{
		IsAuth: true,
		Code:   ErrUnauthorized,
		Msg:    msg,
	}
}
