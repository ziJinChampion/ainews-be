package model

import (
	"encoding/base64"

	"github.com/southwind/ainews/pkg/constant"
)

type User struct {
	Model

	Name     string `json:"name" gorm:"column:user_name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserRole string `json:"user_role" gorm:"column:user_role"`
	Grade    int    `json:"grade"`
}

func GetUserInfo(maps interface{}) (user User, err error) {
	err = client.Table("users").Where(maps).Find(&user).Error
	return
}

func ValidUserInfo(name, password string) (bool, error) {
	var user User
	pass := base64.StdEncoding.EncodeToString([]byte(password))
	err := client.Table("users").Select("id").Where("user_name = ? and password = ?", name, pass).First(&user).Error

	return user.Id > 0, err
}

func RegisterUser(name, password, mobile, email string) (bool, error) {

	pass := base64.StdEncoding.EncodeToString([]byte(password))

	err := client.Table("users").Create(&User{
		Name:     name,
		Password: pass,
		Mobile:   mobile,
		Email:    email,
		UserRole: constant.NORMAL_USER,
		Grade:    1,
	}).Error

	return true, err
}
