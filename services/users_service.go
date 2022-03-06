package services

import (
	"github/atakanteko/bookstore_users-api/domain/users"
	"github/atakanteko/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
