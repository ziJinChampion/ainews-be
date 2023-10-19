package repository

import "github.com/southwind/ainews/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(string) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByNameAndPassword(string, string) (*entity.User, map[string]string)
}
