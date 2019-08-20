package router

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/dependencyinjection"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/handler"
	"go.uber.org/dig"
)

type router struct {
	dig.In

	userHandler handler.UserHandler
}

// NewRouter is call routing function
func NewRouter(e *echo.Echo, container dependencyinjection.Container) {
	var userHandler handler.UserHandlerInterface
	_ = container.GetContainer().Invoke(func(handler handler.UserHandlerInterface) {
		userHandler = handler
	})

	e.GET("/users", userHandler.GetUsers)

	e.GET("/users/1", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
}
