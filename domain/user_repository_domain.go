package domain

import "github.com/google/uuid"

type UserRepositoryFindFirst struct {
	Email     string
	FirstName string
	LastName  string
}

type UserRepositoryFindFirstReturn struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}

type UserRepositoryCreateUserArgs struct {
	Email     string
	FirstName string
	LastName  string
	Password  *string
}

type UserRepositoryCreateUserReturn struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}
