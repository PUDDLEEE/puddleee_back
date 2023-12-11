package jwtAuth

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtAuthService struct {
	secretKey string
}

func (a *JwtAuthService) CreateAccessToken(claims AuthTokenClaims) (*string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := accessToken.SignedString([]byte(a.secretKey))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (a *JwtAuthService) CreateRefreshToken(claims AuthTokenClaims) (*string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := refreshToken.SignedString([]byte(a.secretKey))

	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (a *JwtAuthService) ReassignAccessToken(accessToken string, refreshToken string) (*string, error) {
	ok, err := a.VerifyRefreshToken(refreshToken)
	if err != nil || !ok {
		return nil, err
	}
	parsedAccessToken, err := a.ParseAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	parsedRefreshToken, err := a.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	translatedAccessToken, ok := parsedAccessToken.Claims.(*AuthTokenClaims)
	if !ok {
		internalError := errors.INTERNAL_ERROR
		internalError.Msg = "internal server error while translate Access Token"
		return nil, internalError
	}
	translatedRefreshToken, ok := parsedRefreshToken.Claims.(*AuthTokenClaims)
	if !ok {
		internalError := errors.INTERNAL_ERROR
		internalError.Msg = "internal server error while translate Refresh Token"
		return nil, internalError
	}

	if translatedAccessToken.ID != translatedRefreshToken.ID {
		notAuthorizedError := errors.NOT_AUTHORIZED
		notAuthorizedError.Msg = "You're trying reassign access token with invalid refresh token which is not yours."
		return nil, notAuthorizedError
	}
	newAccessToken, err := a.CreateAccessToken(*translatedAccessToken)
	if err != nil {
		return nil, err
	}
	return newAccessToken, nil
}

func (a *JwtAuthService) ParseAccessToken(accessToken string) (*jwt.Token, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedAccessToken, nil
}

func (a *JwtAuthService) ParseRefreshToken(accessToken string) (*jwt.Token, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(accessToken, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedRefreshToken, nil
}

func (a *JwtAuthService) VerifyAccessToken(accessToken string) (bool, error) {
	token, err := a.ParseAccessToken(accessToken)
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*AuthTokenClaims); ok {
		return true, nil
	}

	return false, err
}

func (a *JwtAuthService) VerifyRefreshToken(refreshToken string) (bool, error) {
	token, err := a.ParseAccessToken(refreshToken)
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*AuthTokenClaims); ok {
		return true, nil
	}
	return false, err
}

func NewJwtAuthService(secretKey string) *JwtAuthService {
	return &JwtAuthService{secretKey: secretKey}
}
