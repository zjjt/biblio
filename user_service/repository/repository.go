package repository

import (
	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"
	userP "github.com/zjjt/biblio/user_service/proto/user"
)

//Repository is an interface
type Repository interface {
	CreateUser(user *userP.User) error
	GetUserByName(name string) (*userP.User, error)
}

//UserRepository implements the methods defined by the Repository interface
type UserRepository struct {
	db *gorm.DB
}

//CreateUser creates a user in database
func (repo *UserRepository) CreateUser(user *userP.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		log.Errorf("[CreateUser repository] couldnt create user in the database %v", err)
		return err
	}
	return nil
}

//GetUserByName get the User based on the name of the user in the DB
func (repo *UserRepository) GetUserByName(name string) (*userP.User, error) {
	user := new(userP.User)
	if err := repo.db.Where("name=?", name).First(&user).Error; err != nil {
		log.Errorf("[GetUserByName repository] couldnt get user from the database %v", err)
		return nil, err
	}
	return user, nil

}

//InitRepository creates an instance of a UserRepository to query the database
func InitRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
