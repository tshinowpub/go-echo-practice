package router

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/provider"
)

func NewRouter(e *echo.Echo, provider provider.ServiceProvider) {
	e.GET("/users", provider.BuildUserHandler().GetUsers)
	e.POST("/users", provider.BuildUserHandler().CreateUser)
}
