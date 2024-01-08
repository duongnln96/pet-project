package serror

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ErrSystemInternal string = "err_system_internal"
	ErrUserCommon     string = "err_user_common"
	ErrNotFound       string = "err_not_found"
)

type EchoErrorResponse struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

func (m *EchoErrorResponse) Error() string {
	bytes, _ := json.Marshal(*m)
	return string(bytes)
}

func NewErrorResponse(httpCode int, errCode string, errMsg string) *EchoErrorResponse {
	return &EchoErrorResponse{
		Code:      httpCode,
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
	}
}

type EchoResponse struct {
	EchoErrorResponse `json:",inline"`
	Data              interface{} `json:"data,omitempty"`
}

func EchoSuccess(data interface{}) (int, EchoResponse) {
	return http.StatusOK, EchoResponse{
		EchoErrorResponse: EchoErrorResponse{
			Code: http.StatusOK,
		},
		Data: data,
	}
}

func CustomEchoErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}

	response := &EchoResponse{
		EchoErrorResponse: EchoErrorResponse{
			Code:     http.StatusBadRequest,
			ErrorMsg: err.Error(),
		},
		Data: nil,
	}

	if he, ok := err.(*echo.HTTPError); ok {
		response.Code = he.Code
		response.ErrorCode = ErrNotFound

		switch m := he.Message.(type) {
		case string:
			response.ErrorMsg = m
		case json.Marshaler:
			// do nothing - this type knows how to format itself to JSON
		case error:
			response.ErrorMsg = m.Error()
		}

	} else if appErr, ok := err.(*EchoErrorResponse); ok {

		response.Code = appErr.Code
		response.ErrorMsg = appErr.ErrorMsg
		response.ErrorCode = appErr.ErrorCode
	}

	if c.Request().Method == echo.HEAD {
		c.NoContent(response.Code)
	} else {
		c.JSON(response.Code, response)
	}
}

func Service2EchoErr(sErr error) error {
	serr, ok := sErr.(*SError)
	if !ok {
		return NewErrorResponse(http.StatusInternalServerError, ErrSystemInternal, "Error occur while handling convert error -> echo error")
	}

	if serr.IsSystem {
		return NewErrorResponse(http.StatusInternalServerError, serr.Code, serr.Msg)
	} else {
		return NewErrorResponse(http.StatusOK, serr.Code, serr.Msg)
	}
}

func HealthHandler(c echo.Context) error {
	return c.JSON(EchoSuccess(nil))
}
