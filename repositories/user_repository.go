package repositories

import (
	"be-test/config"
	"be-test/domain"
	"be-test/ent/predicate"
	"be-test/ent/user"
	"context"
	"errors"
)

type UserRepository interface {
	FindFirst(
		context context.Context,
		args domain.UserRepositoryFindFirst,
	) (
		domain.UserRepositoryFindFirstReturn,
		error,
	)
	Create(
		context context.Context,
		args domain.UserRepositoryCreateUserArgs,
	) (
		domain.UserRepositoryCreateUserReturn,
		error,
	)
}

type UserRepositoryImpl struct{}

func (u *UserRepositoryImpl) FindFirst(
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

func (u *UserRepositoryImpl) Create(
	context context.Context,
	args domain.UserRepositoryCreateUserArgs,
) (
	domain.UserRepositoryCreateUserReturn,
	error,
) {
	tx, err := config.DB.Tx(context)
	if err != nil {
		return domain.UserRepositoryCreateUserReturn{}, err
	}
	defer tx.Rollback()

	userAlreadyExistWithEmail, err := tx.User.Query().Where(user.EmailEQ(args.Email)).Exist(context)
	if err != nil {
		return domain.UserRepositoryCreateUserReturn{}, err
	}

	if userAlreadyExistWithEmail {
		return domain.UserRepositoryCreateUserReturn{}, errors.New("User already exist with email")
	}

	query := tx.User.Create().SetFirstName(args.FirstName).SetLastName(args.LastName).SetEmail(args.Email)

	if args.Password != nil {
		cred, err := tx.Credential.Create().SetPassword(*args.Password).Save(context)
		if err != nil {
			return domain.UserRepositoryCreateUserReturn{}, err
		}
		query.SetCredentialID(cred.ID)
	}

	user, err := query.Save(context)

	if err != nil {
		return domain.UserRepositoryCreateUserReturn{}, err
	}

	if err := tx.Commit(); err != nil {
		return domain.UserRepositoryCreateUserReturn{}, err
	}

	return domain.UserRepositoryCreateUserReturn{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil

}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}
