package usecase

import (
	"context"
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

func (u *userUsecase) CreateUser(ctx context.Context, user entity.User) (err error) {
	err = u.userRepository.Create(
		ctx,
		user,
	)
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) GetUser(ctx context.Context, id string) (user entity.User, err error) {
	user, err = u.userRepository.FindOne(
		ctx,
		id,
	)
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) GetUsers(ctx context.Context) (users []entity.User, err error) {
	users, err = u.userRepository.FindMany(
		ctx,
	)
	if err != nil {
		return
	}

	return users, nil
}
