package auth

import (
	"github.com/crafty-ezhik/auth/internal/cfg"
	"github.com/crafty-ezhik/auth/internal/err"
	"time"
)

// Config - configuration for creating a pair of tokens
type Config struct {
	SigningKey string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

// NewDefaultConfig - returns a pointer to Config with default values: AccessTTL=15m, RefreshTTL=7d.
func NewDefaultConfig(key string) *Config {
	return &Config{
		SigningKey: key,
		AccessTTL:  cfg.DefaultAccessTTL,
		RefreshTTL: cfg.DefaultRefreshTTL,
	}
}

// NewConfig - returns a pointer to the Config with the passed parameters.
//
// There are restrictions, accessTTL should be from 1 minute to 1 hour,
// and refreshTTL should be at least 5 minutes and no more than 30 days.
//
// If the values are incorrect, an error is returned.
func NewConfig(key string, accessTTL, refreshTTL time.Duration) (*Config, error) {
	if accessTTL >= time.Hour && accessTTL < time.Minute {
		return nil, err.ErrAccessTokenLifetime
	}
	if refreshTTL >= time.Hour*720 && refreshTTL < time.Minute*5 {
		return nil, err.ErrRefreshTokenLifetime
	}
	return &Config{
		SigningKey: key,
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}, nil
}
