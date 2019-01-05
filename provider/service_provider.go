package provider

import (
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/handler"
	"github.com/tshinowpub/go-echo-practice/interface/controllers"
)

type serviceProvider struct {
}

type ServiceProvider interface {
	BuildUserHandler() handler.UserHandler
}

func Boot() ServiceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) BuildUserHandler() handler.UserHandler {
	return handler.BuildUserHandler(sp.BuildUserController())
}

func (sp *serviceProvider) BuildUserController() controllers.UserController {
	return controllers.BuildUserController()
}
