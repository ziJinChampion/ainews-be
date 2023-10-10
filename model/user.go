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
	DB.Where(maps).Find(&user)
	return
}

func ValidUserInfo(name, password string) bool {
	var user User
	pass := base64.StdEncoding.EncodeToString([]byte(password))
	DB.Select("id").Where("name = ? and password = ?", name, pass).First(&user)

	return user.ID > 0
}

func RegisterUser(name, password, mobile, email string) bool {

	pass := base64.StdEncoding.EncodeToString([]byte(password))

	DB.Create(&User{
		Name:     name,
		Password: pass,
		Mobile:   mobile,
		Email:    email,
	})

	return true
}
