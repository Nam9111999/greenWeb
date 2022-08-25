package user

import (
	"context"
	"green.env.com/auth/payload"
	"green.env.com/auth/presenter"
)

type IUseCase interface {
	Create(ctx context.Context, req *payload.CreateRequest) (*presenter.CreateResponseWrapper, error)
}
