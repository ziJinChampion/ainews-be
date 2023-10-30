package security

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	lib "github.com/southwind/ainews/common/config"
)

var jwtSecret = []byte(lib.LoadServerConfig().JwtSecret)

type Claims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	jwt.RegisteredClaims
}

func GenerateToken(username, role, email, mobile string, id uint64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
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
