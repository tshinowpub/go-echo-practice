package dependencyinjection

import (
	"github.com/pkg/errors"
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

	/**
	 * controllers
	 */
	if err := digContainer.Provide(func(getUser usecases.GetUser) *controllers.UserController {
		return controllers.NewUserController(getUser)
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
