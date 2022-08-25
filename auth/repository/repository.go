package repository

import (
	"context"
	"gorm.io/gorm"
	"green.env.com/auth/repository/user"
)

type Repository struct {
	User user.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User: user.NewPG(getClient),
	}
}
