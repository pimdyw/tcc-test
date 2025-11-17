package usecase

import (
	"tcc-test/internal/entity"
	"tcc-test/internal/repository"
)

type userUsecase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUsecase(
	userRepository repository.UserRepositoryInterface,
) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CreateUser(user entity.User, err error) {
	panic("unimplemented")
}

func (u *userUsecase) GetUser() (user entity.User, err error) {
	panic("unimplemented")
}

func (u *userUsecase) GetUsers() (user []entity.User, err error) {
	panic("unimplemented")
}
