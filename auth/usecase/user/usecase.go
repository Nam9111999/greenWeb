package user

import (
	"green.env.com/auth/repository"
	"green.env.com/auth/repository/user"
)

type UseCase struct {
	UserRepo user.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}
