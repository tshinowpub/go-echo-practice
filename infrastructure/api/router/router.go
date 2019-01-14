package router

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/provider"
)

func NewRouter(e *echo.Echo, provider provider.ServiceProvider) {
	e.GET("/users", provider.BuildUserHandler().GetUsers)
	e.GET("/users/1", provider.BuildUserHandler().GetUser)
	e.POST("/users", provider.BuildUserHandler().CreateUser)
}
