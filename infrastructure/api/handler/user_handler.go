package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/interface/controllers"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}

func BuildUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUsers")
}

func (uh *userHandler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}
