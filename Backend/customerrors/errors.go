package customerrors

import "net/http"

type CustomError struct {
	ErrorCode    int
	ErrorMessage string
}

func (err CustomError) Error() string {
	return err.ErrorMessage
}

func NewError(errCode int, errMsg string) *CustomError {
	return &CustomError{
		ErrorCode:    errCode,
		ErrorMessage: errMsg,
	}
}

func (err CustomError) GetHTTPErrorCode() int {
	switch err.ErrorCode {
	case InvalidAction:
		return http.StatusBadRequest
	case AlreadyExist:
		return http.StatusBadRequest
	case NotExist:
		return http.StatusNotFound
	case Unauthorized:
		return http.StatusUnauthorized
	case DatabaseError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

const (
	InvalidAction = 40000
	AlreadyExist  = 40001
	NotExist      = 40002
	Unauthorized  = 40003
	CommonErr     = 50000
	DatabaseError = 50001
)
