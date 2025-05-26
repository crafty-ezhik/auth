package auth

import "strings"

type ContextField int

const (
	UserID ContextField = iota
	Username
	Email
	Role
	Permissions
)

type MiddlewareOption func(*middlewareConfig)

type middlewareConfig struct {
	fields []ContextField
}

func WithFields(fields ...ContextField) MiddlewareOption {
	return func(cfg *middlewareConfig) {
		cfg.fields = append(cfg.fields, fields...)
	}
}

func AuthMiddleware(authService Auth, opts ...MiddlewareOption) func(HTTPContext) {

	// Добавляем поля, переданные через функции WithFields
	cfg := &middlewareConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	return func(c HTTPContext) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Unauthorized()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Unauthorized()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		data, err := authService.ParseToken(tokenString)
		if err != nil {
			// TODO: передать err в Unauthorized, чтобы потом поместить в ответ
			c.Unauthorized()
			return
		}
		userID, err := data.GetUserID()
		if err != nil {
			c.Unauthorized()
			return
		}
		c.SetUser(userID)
		c.Next()
	}
}
