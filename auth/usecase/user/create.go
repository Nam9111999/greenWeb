package user

import (
	"context"
	"github.com/pkg/errors"
	"green.env.com/auth/payload"
	"green.env.com/auth/presenter"
	"green.env.com/auth/util"
	"net/http"
	"strings"
)

func (u *UseCase) validateCreate(req *payload.CreateRequest) error {
	req.UserName = strings.TrimSpace(req.UserName)
	if req.UserName == "" {
		// TODO: custom err
		return util.NewError(errors.New("Invalid param"), http.StatusBadRequest, "400", "", false)
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateRequest,
) (*presenter.CreateResponseWrapper, error) {
	return nil, nil
}
