package repositories

import (
	"be-test/config"
	"be-test/domain"
	"be-test/ent/predicate"
	"be-test/ent/user"
	"context"
)

type UserRepository interface {
	FindFirst(
		context context.Context,
		args domain.UserRepositoryFindFirst,
	) (
		domain.UserRepositoryFindFirstReturn,
		error,
	)
}

type UserRepositoryImpl struct{}

func (u UserRepositoryImpl) FindFirst(
	context context.Context,
	args domain.UserRepositoryFindFirst,
) (
	domain.UserRepositoryFindFirstReturn,
	error,
) {
	query := config.DB.User.Query()
	var predicates []predicate.User

	if args.FirstName != "" {
		predicates = append(
			predicates,
			user.FirstNameContains(args.FirstName),
		)
	}

	if args.LastName != "" {
		predicates = append(
			predicates,
			user.LastNameContains(args.LastName),
		)
	}

	if args.Email != "" {
		predicates = append(
			predicates,
			user.EmailContains(args.Email),
		)
	}

	if len(predicates) > 0 {
		query.Where(predicates...)
	}

	data, err := query.First(context)

	if err != nil {
		return domain.UserRepositoryFindFirstReturn{}, err
	}

	return domain.UserRepositoryFindFirstReturn{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}, nil
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}
