package repository

import (
	"context"
	"database/sql"
	"tcc-test/internal/entity"
)

type userRepository struct {
	userTable *sql.DB
}

func NewUserRepository(
	userTable *sql.DB,
) *userRepository {
	return &userRepository{
		userTable: userTable,
	}
}

func (u *userRepository) create(ctx context.Context, user entity.User) (err error) {
	panic("unimplemented")
}

func (u *userRepository) findMany(ctx context.Context, id string) (user entity.User, err error) {
	panic("unimplemented")
}

func (u *userRepository) findOne(ctx context.Context) (users []entity.User, err error) {
	panic("unimplemented")
}
