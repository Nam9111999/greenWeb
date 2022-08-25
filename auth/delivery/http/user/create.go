package user

import (
	"github.com/labstack/echo/v4"
	"green.env.com/auth/payload"
	"green.env.com/auth/presenter"
	"green.env.com/auth/util"
)

func (r *Route) Create(c echo.Context) error {
	var (
		ctx = util.CustomEchoContext{
			Context: c,
		}
		req  = &payload.CreateRequest{}
		resp *presenter.CreateResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return util.Response.Error(ctx, util.ErrJSONMarshal(err))
	}

	resp, err := r.UseCase.User.Create(&ctx, req)
	if err != nil {
		return util.Response.Error(c, err.(util.CustomError))
	}

	return util.Response.Success(c, resp)
}
