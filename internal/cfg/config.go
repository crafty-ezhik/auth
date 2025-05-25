package cfg

import "time"

// InternalConfig - internal configuration for creating a pair of tokens
type InternalConfig struct {
	SigningKey string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}
