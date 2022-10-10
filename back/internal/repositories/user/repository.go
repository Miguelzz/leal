package user_repo

import (
	"loyalty/internal/core/domain"
)

var users = []domain.User{
	{
		ID:       1,
		IdPoints: 1,
	},
}

type userTable struct {
}

func New() *userTable {
	return &userTable{}
}

func (repo *userTable) Get(id int) (domain.User, error) {
	return users[0], nil
}

func (repo *userTable) Update(_user domain.User) (domain.User, error) {
	users[_user.ID-1] = _user
	return _user, nil
}
