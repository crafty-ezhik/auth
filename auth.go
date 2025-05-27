package auth

import (
	"github.com/crafty-ezhik/auth/internal/cfg"
	"github.com/crafty-ezhik/auth/internal/service"
	"github.com/crafty-ezhik/auth/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct {
	cfg *cfg.InternalConfig
}

func NewAuth(config *Config) *Auth {
	return &Auth{
		cfg: &cfg.InternalConfig{
			SigningKey: config.SigningKey,
			AccessTTL:  config.AccessTTL,
			RefreshTTL: config.RefreshTTL,
		},
	}
}

func (auth *Auth) GenerateAccessToken(claims jwt.MapClaims) (types.AccessToken, error) {
	claims["exp"] = time.Now().Add(auth.cfg.AccessTTL).Unix()
	tokenClaims := &types.CustomClaims{
		MapClaims: claims,
	}
	accessToken, err := service.GenerateAccessToken(auth.cfg.SigningKey, tokenClaims)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (auth *Auth) GenerateRefreshToken(claims jwt.MapClaims) (types.RefreshToken, error) {
	claims["exp"] = time.Now().Add(auth.cfg.RefreshTTL).Unix()
	tokenClaims := &types.CustomClaims{
		MapClaims: claims,
	}
	refreshToken, err := service.GenerateRefreshToken(auth.cfg.SigningKey, tokenClaims)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (auth *Auth) ParseToken(token string) (*types.CustomClaims, error) {
	data, err := service.ParseToken(auth.cfg.SigningKey, token)
	if err != nil {
		return nil, err
	}
	return data, nil
}
