package err

import "errors"

var (
	ErrAccessTokenLifetime  = errors.New("accessTTL must be less than or equal to hour")
	ErrRefreshTokenLifetime = errors.New("refreshTTL must be less than or equal to 30 days")
)
