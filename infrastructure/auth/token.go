package auth

import (
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/southwind/ainews/common/config"
)

type Token struct {
	JwtSecret string
}

func NewToken(config *config.ServerConfig) *Token {
	return &Token{
		JwtSecret: config.JwtSecret,
	}
}

type TokenInterface interface {
	CreateToken(username, role, email, mobile string, id int) (string, error)
	ParseToken(token string) (*UserDetails, error)
}

var _ TokenInterface = &Token{}

func (t *Token) CreateToken(username, role, email, mobile string, id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := UserDetails{
		id,
		username,
		role,
		email,
		mobile,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "ai-news",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(t.JwtSecret)
	return token, err
}

func (t *Token) ParseToken(token string) (*UserDetails, error) {
	token = strings.Fields(token)[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &UserDetails{}, func(token *jwt.Token) (interface{}, error) {
		return t.JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserDetails); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
