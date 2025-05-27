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
			c.Unauthorized("Invalid token")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Unauthorized("Invalid token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		data, err := authService.ParseToken(tokenString)
		if err != nil {
			c.Unauthorized("Invalid token")
			return
		}

		// Добавление необходимых полей в контекст
		for _, field := range cfg.fields {
			switch field {
			case UserID:
				userID, err := data.GetUserID()
				if err != nil {
					c.Unauthorized("userId cannot be added to the context")
					return
				}
				c.SetValueIntoContext("user_id", userID)
			case Username:
				username, err := data.GetUsername()
				if err != nil {
					c.Unauthorized("username cannot be added to the context")
					return
				}
				c.SetValueIntoContext("username", username)
			case Email:
				email, err := data.GetEmail()
				if err != nil {
					c.Unauthorized("email cannot be added to the context")
					return
				}
				c.SetValueIntoContext("email", email)
			case Role:
				role, err := data.GetRole()
				if err != nil {
					c.Unauthorized("role cannot be added to the context")
					return
				}
				c.SetValueIntoContext("role", role)
			case Permissions:
				permissions, err := data.GetPermissions()
				if err != nil {
					c.Unauthorized("permissions cannot be added to the context")
					return
				}
				c.SetValueIntoContext("permissions", permissions)
			}
		}
		c.Next()
	}
}
