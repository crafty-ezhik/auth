package auth

import "github.com/crafty-ezhik/auth/internal/cfg"

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

}
