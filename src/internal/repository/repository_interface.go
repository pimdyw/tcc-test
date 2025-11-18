package repository

import (
	"context"
	"tcc-test/internal/entity"
)

type UserRepositoryInterface interface {
	Create(
		ctx context.Context,
		user entity.User,
	) (err error)

	FindOne(
		ctx context.Context,
		id string,
	) (user entity.User, err error)

	FindMany(
		ctx context.Context,
	) (user []entity.User, err error)
}
