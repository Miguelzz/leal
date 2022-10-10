package user_repo

import (
	"loyalty/internal/core/domain"
	"loyalty/internal/repositories/mocks"
)

type userTable struct {
}

func New() *userTable {
	return &userTable{}
}

func (repo *userTable) Get(id int) (domain.User, error) {
	return mocks.MockUsers[0], nil
}
