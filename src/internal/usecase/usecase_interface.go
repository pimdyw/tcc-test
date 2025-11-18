package usecase

import (
	"context"
	"tcc-test/internal/entity"
)

type UserUsecaseInterface interface {
	CreateUser(ctx context.Context, user entity.User) (err error)
	GetUsers(ctx context.Context,) (user []entity.User, err error)
	GetUser(ctx context.Context, id string) (user entity.User, err error)
}
