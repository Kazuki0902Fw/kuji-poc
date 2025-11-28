package authtoken

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"kujicole/domain/model"
	"kujicole/domain/service"
	"kujicole/infra/bob/models"
	"kujicole/util"
	"github.com/aarondl/opt/omit"
	"github.com/pkg/errors"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
)

type authTokenService struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	conn       *bob.DB
}

var _ service.AuthTokenService = &authTokenService{}

func NewAuthTokenService(privateKeyBase64 string, publicKeyBase64 string, conn *bob.DB) (*authTokenService, error) {
	prvByte, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	prv, err := x509.ParsePKCS1PrivateKey(prvByte)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pubByte, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pub, err := x509.ParsePKCS1PublicKey(pubByte)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &authTokenService{privateKey: prv, publicKey: pub, conn: conn}, nil
}

func (s *authTokenService) Create(ctx context.Context, userID model.ID) (*model.AuthToken, error) {
	now := time.Now()

	atExp, err := tokenTypeAccessToken.GetExpirationTime(now)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claimsForAccessToken := &customClaims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: atExp,
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   string(tokenTypeAccessToken),
		},
	}
	atToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claimsForAccessToken)
	atTokenString, err := atToken.SignedString(s.privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rfExp, err := tokenTypeRefreshToken.GetExpirationTime(now)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claimsForRefreshToken := &customClaims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: rfExp,
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   string(tokenTypeRefreshToken),
		},
	}
	rfToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claimsForRefreshToken)
	rfTokenString, err := rfToken.SignedString(s.privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	exists, err := models.RefreshTokens.Query(
		sm.Where(models.RefreshTokens.Columns.UserID.EQ(mysql.Arg(userID.String()))),
	).Exists(ctx, s.conn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if exists {
		refreshTokenRecord, err := models.RefreshTokens.Query(
			sm.Where(models.RefreshTokens.Columns.UserID.EQ(mysql.Arg(userID.String()))),
		).One(ctx, s.conn)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		currentRefreshTokenSetter := &models.RefreshTokenSetter{
			ExpiresAt: omit.From(rfExp.Time),
			TokenHash: omit.From(util.HashSHA256(rfTokenString)),
		}
		err = refreshTokenRecord.Update(ctx, s.conn, currentRefreshTokenSetter)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		refreshTokenRecordSetter := &models.RefreshTokenSetter{
			ID:        omit.From(model.NewID().String()),
			UserID:    omit.From(userID.String()),
			TokenHash: omit.From(util.HashSHA256(rfTokenString)),
			ExpiresAt: omit.From(rfExp.Time),
		}
		_, err = models.RefreshTokens.Insert(refreshTokenRecordSetter).One(ctx, s.conn)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &model.AuthToken{
		AccessToken:           atTokenString,
		AccessTokenExpiresAt:  atExp.Time,
		RefreshToken:          rfTokenString,
		RefreshTokenExpiresAt: rfExp.Time,
	}, nil
}

func (s *authTokenService) UpdateByRefreshToken(ctx context.Context, refreshToken string) (*model.AuthToken, error) {
	claims, err := s.verifyToken(ctx, refreshToken, tokenTypeRefreshToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hashedRefreshToken := util.HashSHA256(refreshToken)
	currentRefreshToken, err := models.RefreshTokens.Query(
		sm.Where(models.RefreshTokens.Columns.TokenHash.EQ(mysql.Arg(hashedRefreshToken))),
		sm.Where(models.RefreshTokens.Columns.UserID.EQ(mysql.Arg(claims.UserID))),
		sm.Where(models.RefreshTokens.Columns.ExpiresAt.EQ(mysql.Arg(claims.ExpiresAt.Time))),
	).One(ctx, s.conn)
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(fmt.Errorf("refresh token not found. token_hash: %s, user_id: %s, expires_at: %s", hashedRefreshToken, claims.UserID, claims.ExpiresAt.Time))
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := model.IDFromString(claims.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	newToken, err := s.Create(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	newRefreshTokenHash, err := s.hashRefreshToken(newToken.RefreshToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	currentRefreshTokenSetter := &models.RefreshTokenSetter{
		ExpiresAt: omit.From(newToken.RefreshTokenExpiresAt),
		TokenHash: omit.From(newRefreshTokenHash),
	}
	err = currentRefreshToken.Update(ctx, s.conn, currentRefreshTokenSetter)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return newToken, nil
}

func (s *authTokenService) VerifyAccessToken(ctx context.Context, token string) (*model.ID, error) {
	claims, err := s.verifyToken(ctx, token, tokenTypeAccessToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := model.IDFromString(claims.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &userID, nil
}

func (s *authTokenService) DeleteRefreshTokenByUserID(ctx context.Context, userID model.ID) error {
	records, err := models.RefreshTokens.Query(
		sm.Where(models.RefreshTokens.Columns.UserID.EQ(mysql.Arg(userID.String()))),
	).All(ctx, s.conn)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, record := range records {	
		err = record.Delete(ctx, s.conn)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (s *authTokenService) verifyToken(ctx context.Context, token string, tokenType tokenType) (*customClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !parsedToken.Valid {
		return nil, errors.WithStack(errors.New("invalid token"))
	}

	claims, ok := parsedToken.Claims.(*customClaims)
	if !ok {
		return nil, errors.WithStack(errors.New("invalid token"))
	}
	if claims.GetSubject() != string(tokenType) {
		return nil, errors.WithStack(fmt.Errorf("expected token type: %s. invalid token type in claims sub: %s", tokenType, claims.GetSubject()))
	}
	now := time.Now()
	if claims.ExpiresAt.Time.Before(now) {
		return nil, errors.WithStack(errors.New("token expired"))
	}

	return claims, nil
}

func (s *authTokenService) hashRefreshToken(token string) (string, error) {
	h := sha256.New()
	_, err := io.WriteString(h, token)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
