package services

import (
	"be-test/domain"
	"be-test/repositories"
	"context"
)

type UserService interface {
	FindUser(
		context context.Context,
		arg domain.UserServiceFindUserArg,
	) string
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func (u *UserServiceImpl) FindUser(
	context context.Context,
	arg domain.UserServiceFindUserArg,
) string {
	//TODO implement me
	data, err := u.userRepository.FindFirst(
		context,
		domain.UserRepositoryFindFirst{
			Email:     arg.Email,
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
		},
	)
	if err != nil {
		return ""
	}

	return data.ID.String()
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}
