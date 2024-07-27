package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

// GenToken 生成token
func GenToken(secret, issuer string, expireHour, userId int, userName string) (string, error) {
	claims := MyClaims{
		Id:       userId,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析token
func ParseToken(secret, token string) (*MyClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		switch vError, ok := err.(*jwt.ValidationError); ok {
		case vError.Errors&jwt.ValidationErrorMalformed != 0:
			return nil, ErrTokenMalformed
		case vError.Errors&jwt.ValidationErrorExpired != 0:
			return nil, ErrTokenExpired
		case vError.Errors&jwt.ValidationErrorNotValidYet != 0:
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}

	if claims, ok := jwtToken.Claims.(*MyClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
