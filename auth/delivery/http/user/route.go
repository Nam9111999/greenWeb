package user

import (
	"github.com/labstack/echo/v4"
	"green.env.com/auth/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.GET("", r.Create)
}
