package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/interface/controllers"
)

type UserHandler struct {
	userController controllers.UserControllerInterface
}

type UserHandlerInterface interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

func NewUserHandler(userController controllers.UserControllerInterface) UserHandlerInterface {
	return &UserHandler{userController}
}

func (uh *UserHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUsers")
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	user, err := uh.userController.GetUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}
