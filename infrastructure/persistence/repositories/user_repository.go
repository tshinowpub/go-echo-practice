package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tshinowpub/go-echo-practice/domain/entities"
	repositoryInterfaces "github.com/tshinowpub/go-echo-practice/domain/repositories"
	"github.com/tshinowpub/go-echo-practice/infrastructure/persistence/models"
)

type UserRepository struct {
	db *gorm.DB
	gormUser models.GormUser
}

func NewUserRepository(db *gorm.DB, gormUser models.GormUser) repositoryInterfaces.UserRepositoryInterface {
	return &UserRepository{db: db, gormUser: gormUser}
}

func TestUserRepository(db *gorm.DB, gormUser models.GormUser) *UserRepository {
	return &UserRepository{db: db, gormUser: gormUser}
}

func (userRepository *UserRepository) Add(user entities.User) (entities.User, error) {
	gormUsers := &models.GormUser{
		UserAccountId: user.UserAccountId,
		UserName: user.UserName,
		Email: user.Email,
	}

	userRepository.db.Clone();

	fmt.Printf("%#v\n", userRepository.db)
	fmt.Printf("%#v\n", gormUsers)

	userRepository.db.Create(gormUsers)
	fmt.Println("Finished !!")

	return user, nil
}