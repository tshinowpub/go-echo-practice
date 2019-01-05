package controllers

type userController struct {
}

type UserController interface {
	GetUsers()
	CreateUsers()
}

func BuildUserController() UserController {
	return &userController{}
}

func (userController *userController) GetUsers() {
}

func (userController *userController) CreateUsers() {
}
