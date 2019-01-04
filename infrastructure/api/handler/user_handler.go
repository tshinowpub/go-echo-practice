package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type userHandler struct {
}

type UserHandler interface {
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}

func CreateUserHandler() UserHandler {
	return &userHandler{}
}

func (uh *userHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUsers")
}

func (uh *userHandler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}
