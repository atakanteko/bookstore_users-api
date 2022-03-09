package users

import (
	"fmt"
	"github/atakanteko/bookstore_users-api/utils/date_utils"
	"github/atakanteko/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
func (user *User) Save() *errors.RestErr {
	result := userDB[user.ID]
	if result != nil {
		if result.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))

		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.ID] = user
	return nil
}
