package router

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/dependencyinjection"
	"github.com/tshinowpub/go-echo-practice/interface/controllers"
	"go.uber.org/dig"
)

type router struct {
	dig.In
}

// NewRouter is call routing function
func NewRouter(e *echo.Echo, container dependencyinjection.Container) {
	var userController controllers.UserController
	_ = container.GetContainer().Invoke(func(controller controllers.UserController) {
		userController = controller
	})

	e.GET("/users", userController.GetUsers)

	e.GET("/users/1", userController.GetUser)
	e.POST("/users", userController.CreateUser)
}
