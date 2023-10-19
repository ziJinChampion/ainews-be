package persistence

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
	"github.com/southwind/ainews/utils/security"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

var _ repository.UserRepository = &UserDAO{}

func (r *UserDAO) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserDAO) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Create(&user).Error
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

func (r *UserDAO) GetUser(name string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDAO) GetUserByNameAndPassword(name, password string) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	err := r.db.Where("name = ?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil

}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db}
}
