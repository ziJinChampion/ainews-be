package application

import (
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
)

type userApp struct {
	userRepo repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
	GetUser(string) (*entity.User, error)
	GetUserByNameAndPassword(name, passwod string) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.userRepo.SaveUser(user)
}

func (u *userApp) GetUser(name string) (*entity.User, error) {
	return u.userRepo.GetUser(name)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.userRepo.GetUsers()
}

func (u *userApp) GetUserByNameAndPassword(name, password string) (*entity.User, map[string]string) {
	return u.userRepo.GetUserByNameAndPassword(name, password)
}
