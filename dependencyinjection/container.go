package dependencyinjection

import (
	"github.com/pkg/errors"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/handler"
	"github.com/tshinowpub/go-echo-practice/interface/controllers"
	"github.com/tshinowpub/go-echo-practice/usecase/usecases"
	"go.uber.org/dig"
)

type container struct {
	container *dig.Container
}

// Container is managed dependency injection in application
type Container interface {
	GetContainer() *dig.Container
}

// New Initialize Di container
func New() (Container, error)  {
	digContainer := dig.New()

	if err := digContainer.Provide(func(getUser usecases.GetUser) controllers.UserControllerInterface {
		return controllers.NewUserController(getUser)
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	/**
	 * Handlers
	 */
	if err := digContainer.Provide(func(userController controllers.UserControllerInterface) handler.UserHandlerInterface {
		return handler.NewUserHandler(userController)
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	/**
	 * Usecases
	 */
	if err :=  digContainer.Provide(func() usecases.GetUser {
		return usecases.GetUser{}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &container{container: digContainer}, nil
}

// Get Di container
func (container *container) GetContainer() *dig.Container {
	return container.container
}
