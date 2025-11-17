package repository

import (
	"context"
	"tcc-test/internal/entity"
)

type UserRepositoryInterface interface {
	create(
		ctx context.Context,
		user entity.User,
	) (err error)

	findOne(
		ctx context.Context,
	) (users []entity.User, err error)

	findMany(
		ctx context.Context,
		id string,
	) (user entity.User, err error)
}
