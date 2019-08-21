package controllers

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/usecase/usecases"
	"net/http"
)

type UserController struct {
	getUser usecases.GetUser
}

func NewUserController(getUser usecases.GetUser) *UserController {
	return &UserController{getUser: getUser}
}

func (userController *UserController) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUsers")
}

func (userController *UserController) GetUser(c echo.Context) error {
	user, err := userController.getUser.Run()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (userController *UserController) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}
