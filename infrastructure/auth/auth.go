package auth

import jwt "github.com/golang-jwt/jwt/v5"

type AuthInterface interface {
	CreateAuth()
}

type UserDetails struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	jwt.RegisteredClaims
}
