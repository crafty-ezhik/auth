package service

import (
	"fmt"
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

func ParseToken(singingKey, token string) (*types.CustomClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(singingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	data, ok := parsedToken.Claims.(*types.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type")
	}
	return data, nil
}

func generateToken(singingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString([]byte(singingKey))
	if err != nil {
		return "", err
	}
	return signToken, nil
}
