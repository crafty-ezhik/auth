package service

import (
	"github.com/crafty-ezhik/auth/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(singingKey string, claims jwt.Claims) (types.AccessToken, error) {
	token, err := generateToken(singingKey, claims)
	if err != nil {
		return "", err
	}
	return types.AccessToken(token), nil
}

func GenerateRefreshToken(singingKey string, claims jwt.Claims) (types.RefreshToken, error) {
	token, err := generateToken(singingKey, claims)
	if err != nil {
		return "", err
	}
	return types.RefreshToken(token), nil
}

func Verify() bool {
	return true
}

func generateToken(singingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString([]byte(singingKey))
	if err != nil {
		return "", err
	}
	return signToken, nil
}
