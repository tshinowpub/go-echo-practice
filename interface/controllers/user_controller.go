package controllers

import (
	"github.com/tshinowpub/go-echo-practice/domain/entities"
)

type userController struct {
}

type UserController interface {
	GetUsers()
	GetUser() (*entities.User, error)
	CreateUsers()
}

func BuildUserController() UserController {
	return &userController{}
}

func (userController *userController) GetUsers() {
}

func (userController *userController) GetUser() (*entities.User, error) {
	user := &entities.User{}

	return user, nil
}

func (userController *userController) CreateUsers() {
}
