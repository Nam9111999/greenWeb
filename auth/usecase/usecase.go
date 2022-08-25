package usecase

import (
	"green.env.com/auth/repository"
	"green.env.com/auth/usecase/user"
)

type UseCase struct {
	User user.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
	}
}
