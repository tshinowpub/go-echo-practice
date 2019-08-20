package controllers

import (
	"github.com/tshinowpub/go-echo-practice/domain/entities"
	"github.com/tshinowpub/go-echo-practice/usecase/usecases"
)

type UserController struct {
	getUser usecases.GetUser
}

type UserControllerInterface interface {
	GetUsers()
	GetUser() (*entities.User, error)
	CreateUsers()
}

func NewUserController(getUser usecases.GetUser) UserControllerInterface {
	return &UserController{getUser: getUser}
}

func (userController *UserController) GetUsers() {
}

func (userController *UserController) GetUser() (*entities.User, error) {
	user, err := userController.getUser.Run()
	if err != nil {
		return nil, err
	}

	return user, err
}

func (userController *UserController) CreateUsers() {
}
