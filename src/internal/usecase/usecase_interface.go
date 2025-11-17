package usecase

import "tcc-test/internal/entity"

type UserUsecaseInterface interface {
	CreateUser(user entity.User, err error)
	GetUser() (user entity.User, err error)
	GetUsers() (user []entity.User, err error)
}
