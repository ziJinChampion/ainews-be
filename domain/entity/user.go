package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/southwind/ainews/common/constant"
	"github.com/southwind/ainews/utils/security"
	"gorm.io/gorm"
)

type User struct {
	Id        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size 100 not null"`
	Password  string    `json:"password" gorm:"size 100 not null"`
	Mobile    string    `json:"mobile" gorm:"size 100"`
	Email     string    `json:"email" gorm:"size 100"`
	Role      string    `json:"role" gorm:"size 100;default:'NORMAL_USER'"`
	Grade     int       `json:"grade" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type PublicUser struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Grade int    `json:"grade"`
}

func (u *User) PublicUser() *PublicUser {
	return &PublicUser{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
		Grade: u.Grade,
	}
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) VaildRegisterInfo() (bool, error) {
	if len(u.Name) < 8 || len(u.Name) > 30 {
		return false, errors.New("userName must over than 8 character and less than 30 character")
	}
	if len(u.Password) < 8 || len(u.Password) > 30 {
		return false, errors.New("password must over than 8 character and less than 30 character")
	}
	return true, nil
}

func (u *User) ValidLoginInfo() (bool, error) {
	if len(u.Name) < 8 || len(u.Name) > 30 {
		return false, errors.New("userName must over than 8 character and less than 30 character")
	}
	if len(u.Password) < 8 || len(u.Password) > 30 {
		return false, errors.New("password must over than 8 character and less than 30 character")
	}
	return true, nil
}

func (u *User) PrepareCreate() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Role = constant.NORMAL_USER
}
