package authtoken

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type tokenType string

const (
	tokenTypeAccessToken  tokenType = "access_token"
	tokenTypeRefreshToken tokenType = "refresh_token"

	accessTokenExpirationTime  = time.Hour
	refreshTokenExpirationTime = 24 * time.Hour

	// ローカル環境ではデバッグしやすさのためにアクセストークンの有効期間を長くする
	accessTokenExpirationTimeLocal = time.Hour * 60
)

func (t tokenType) String() string {
	return string(t)
}

func (t tokenType) IsValid() bool {
	return t == tokenTypeAccessToken || t == tokenTypeRefreshToken
}

func (t tokenType) GetExpirationTime(src time.Time) (*jwt.NumericDate, error) {
	switch t {
	case tokenTypeAccessToken:
		if os.Getenv("ENV") == "local" {
			exp := src.Add(accessTokenExpirationTimeLocal)
			return jwt.NewNumericDate(exp), nil
		}
		exp := src.Add(accessTokenExpirationTime)
		return jwt.NewNumericDate(exp), nil
	case tokenTypeRefreshToken:
		exp := src.Add(refreshTokenExpirationTime)
		return jwt.NewNumericDate(exp), nil
	default:
		return nil, errors.WithStack(fmt.Errorf("invalid token type: %s", t))
	}
}
