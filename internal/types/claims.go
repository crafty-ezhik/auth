package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim interface {
	GetUserID() (string, error)
	GetUsername() (string, error)
	GetEmail() (string, error)
	GetRole() (jwt.ClaimStrings, error)
	GetPermissions() (jwt.ClaimStrings, error)
	GetVersion() (uint, error)
	IsValidVersion(currentVersion uint) (bool, error)
}

// CustomClaims - is a type inherited from jwt.MapClaims, but adds new methods to get the most popular fields in JSON.
// The parsing methods are also taken from the jwt library.
type CustomClaims struct {
	jwt.MapClaims
}

func (c *CustomClaims) GetUserID() (string, error) { return c.parseString("user_id") }

func (c *CustomClaims) GetUsername() (string, error) { return c.parseString("username") }

func (c *CustomClaims) GetEmail() (string, error) { return c.parseString("email") }

func (c *CustomClaims) GetRole() (jwt.ClaimStrings, error) { return c.parseSliceString("role") }

func (c *CustomClaims) GetPermissions() (jwt.ClaimStrings, error) {
	return c.parseSliceString("permissions")
}

func (c *CustomClaims) GetVersion() (uint, error) {
	return c.parseInt("version")
}

func (c *CustomClaims) IsValidVersion(currentVersion uint) (bool, error) {
	tVersion, err := c.GetVersion()
	if err != nil {
		return false, err
	}
	if tVersion < currentVersion {
		return false, nil
	}
	return true, nil
}

func (c *CustomClaims) parseString(key string) (string, error) {
	var (
		ok  bool
		raw interface{}
		iss string
	)
	raw, ok = c.MapClaims[key]
	if !ok {
		return "", nil
	}

	iss, ok = raw.(string)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	return iss, nil
}

func (c *CustomClaims) parseInt(key string) (uint, error) {
	var (
		ok  bool
		raw interface{}
		iss uint
	)
	raw, ok = c.MapClaims[key]
	if !ok {
		return 0, nil
	}

	iss, ok = raw.(uint)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}

	return iss, nil
}

func (c *CustomClaims) parseSliceString(key string) (jwt.ClaimStrings, error) {
	var cs []string
	switch v := c.MapClaims[key].(type) {
	case string:
		cs = append(cs, v)
	case []string:
		cs = v
	case []interface{}:
		for _, a := range v {
			vs, ok := a.(string)
			if !ok {
				return nil, jwt.ErrInvalidType
			}
			cs = append(cs, vs)
		}
	}

	return cs, nil
}
