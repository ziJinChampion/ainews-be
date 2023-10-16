package utils

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/southwind/ainews/lib"
)

var jwtSecret = []byte(lib.LoadServerConfig().JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "ai-news",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func (c *Claims) Valid() error {
	if time.Now().After(c.ExpiresAt.Time) {
		return errors.New("this token has expired")
	}
	return nil
}

func ParseToken(token string) (*Claims, error) {
	token = strings.Fields(token)[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
