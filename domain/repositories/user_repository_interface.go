package repositories

import "github.com/tshinowpub/go-echo-practice/domain/entities"

type UserRepositoryInterface interface {
	Add(user entities.User) (entities.User, error)
}


