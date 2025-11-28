package authtoken

import "github.com/golang-jwt/jwt/v4"

type customClaims struct {
	UserID string `json:"user_id"`

	jwt.RegisteredClaims
}

func (c *customClaims) GetSubject() string {
	return c.Subject
}
