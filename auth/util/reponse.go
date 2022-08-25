package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct{}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    data,
	})
}

func (response) Error(c echo.Context, err CustomError) error {
	var errMessage string

	if err.IsSentry {
		GetLogger().Error(err.Raw.Error())
	}

	if err.Raw != nil {
		errMessage = err.Raw.Error()
	}

	return c.JSON(err.HTTPCode, map[string]interface{}{
		"code":    err.ErrorCode,
		"message": err.Message,
		"info":    errMessage,
	})
}
