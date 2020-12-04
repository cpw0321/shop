// Copyright 2020 The shop Authors

// Package utils implements utils.
package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// KEY ...
	KEY string = "JWT-TOKEN"
	// ExprieTime 过期时间30minutes
	ExprieTime int = 60 * 24
	// ISSUER 颁发者
	ISSUER = "shop"
)

var (
	// TokenExpired token过期
	TokenExpired error = errors.New("token is expired")
	// TokenNotValidYet ...
	TokenNotValidYet error = errors.New("token not active yet")
	// TokenMalformed ...
	TokenMalformed error = errors.New("that's not even a token")
	// TokenInvalid ...
	TokenInvalid error = errors.New("couldn't handle this token")
)

// CustomClaims token消息体
type CustomClaims struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 创建token
func CreateToken(userID int, username string) (string, error) {
	claims := CustomClaims{
		userID,
		username,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),                                              // 签发时间
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(ExprieTime)).Unix(), // 过期时间
			Issuer:    ISSUER,                                                         // 签发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(KEY))
	return tokenStr, err
}

// ParseToken 解析token
func ParseToken(tokenstr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenstr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// RefreshToken 刷新token
func RefreshToken(tokenStr string) (string, error) {
	token, err := ParseToken(tokenStr)
	expireAt := time.Now().Add(time.Second * time.Duration(ExprieTime)).Unix()
	newClaims := CustomClaims{
		token.UserID,
		token.Username,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt,
			Issuer:    ISSUER,
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := newToken.SignedString([]byte(KEY))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
