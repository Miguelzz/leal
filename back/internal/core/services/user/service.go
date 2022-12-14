package user_srv

import (
	"errors"
	"loyalty/internal/core/domain"
	"loyalty/internal/core/ports"
)

type service struct {
	userRepository ports.UserRepository
}

func New(userRepository ports.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (srv *service) Get(id int) (domain.User, error) {
	user, err := srv.userRepository.Get(id)
	if err != nil {
		return domain.User{}, errors.New("get user from repository has failed")
	}

	return user, nil
}
