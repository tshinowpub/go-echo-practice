package usecases

import (
	"github.com/tshinowpub/go-echo-practice/domain/entities"
	"github.com/tshinowpub/go-echo-practice/domain/repositories"
)

type GetUser struct {
	userRepository repositories.UserRepository
}

func NewGetUser(userRepository repositories.UserRepository) *GetUser {
	return &GetUser{userRepository: userRepository}
}

func (getUser *GetUser) Run() (*entities.User, error) {
	user := &entities.User{}

	return user, nil
}
