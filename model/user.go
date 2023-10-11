package model

import "encoding/base64"

type User struct {
	Model

	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func GetUserInfo(maps interface{}) (user User) {
	client.Table("users").Where(maps).Find(&user)
	return
}

func ValidUserInfo(name, password string) bool {
	var user User
	pass := base64.StdEncoding.EncodeToString([]byte(password))
	client.Table("users").Select("id").Where("name = ? and password = ?", name, pass).First(&user)

	return user.Id > 0
}

func RegisterUser(name, password, mobile, email string) bool {

	pass := base64.StdEncoding.EncodeToString([]byte(password))

	client.Table("users").Create(&User{
		Name:     name,
		Password: pass,
		Mobile:   mobile,
		Email:    email,
	})

	return true
}
