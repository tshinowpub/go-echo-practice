package persistence

type userRepository struct {
}

type UserRepository interface {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
