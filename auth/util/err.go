package util

import (
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type CustomError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
	IsSentry  bool
}

func (e CustomError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

func (e CustomError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string, isSentry bool) CustomError {
	return CustomError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
		IsSentry:  isSentry,
	}
}

func ErrSendEmail(err error) CustomError {
	return CustomError{
		Raw:       err,
		HTTPCode:  http.StatusNotAcceptable,
		ErrorCode: "001",
		Message:   err.Error(),
		IsSentry:  true,
	}
}

func ErrCommitTransaction(err error) CustomError {
	return CustomError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "002",
		Message:   "Failed to commit transaction",
		IsSentry:  true,
	}
}

func ErrJSONMarshal(err error) CustomError {
	return CustomError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "003",
		Message:   "Failed to json marshal",
		IsSentry:  true,
	}
}

func ErrJSONUnmarshal(err error) CustomError {
	return CustomError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "003",
		Message:   "Failed to json unmarshal",
		IsSentry:  true,
	}
}
