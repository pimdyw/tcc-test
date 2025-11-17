package usecase

import (
	"tcc-test/internal/entity"
	"tcc-test/internal/repository"
	"time"
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

func (u *userUsecase) GetUsers() (users []entity.User, err error) {
	users = []entity.User{
		{
			ID:        "1",
			Name:      "pim",
			Address:   "test",
			BirthDate: time.Time{},
			Age:       25,
		},
	}

	return users, nil
}
