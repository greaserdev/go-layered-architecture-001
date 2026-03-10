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
