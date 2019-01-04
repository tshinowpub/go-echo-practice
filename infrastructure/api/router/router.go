package router

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/handler"
)

func CreateRouter(e *echo.Echo) {
	e.GET("/users", handler.CreateUserHandler().GetUsers)
	e.POST("/users", handler.CreateUserHandler().CreateUser)
}
