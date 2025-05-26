package auth

import (
	"github.com/crafty-ezhik/auth/internal/cfg"
	"github.com/crafty-ezhik/auth/internal/service"
	"github.com/crafty-ezhik/auth/internal/types"
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

func (auth *Auth) GenerateAccessToken() (bool, error) {
	return true, nil
}

func (auth *Auth) GenerateRefreshToken() (bool, error) {
	return true, nil
}

func (auth *Auth) RefreshAccessToken() (bool, error) {
	return true, nil
}

func (auth *Auth) RefreshRequest() (bool, error) {
	return true, nil
}

func (auth *Auth) ParseToken(token string) (*types.CustomClaims, error) {
	data, err := service.ParseToken("", "")
	if err != nil {
		return nil, err
	}

	return nil, nil
}
