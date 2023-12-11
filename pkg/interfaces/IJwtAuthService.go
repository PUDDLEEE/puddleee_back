package interfaces

import (
	"github.com/PUDDLEEE/puddleee_back/internal/jwtAuth"
	"github.com/golang-jwt/jwt/v5"
)

type IJwtAuthService interface {
	CreateAccessToken(jwtAuth.AuthTokenClaims) (*string, error)
	CreateRefreshToken(jwt.RegisteredClaims) (*string, error)
	ReassignAccessToken(string, string) (*string, error)
	ParseAccessToken(string) (*jwt.Token, error)
	ParseRefreshToken(string) (*jwt.Token, error)
	VerifyAccessToken(string) (bool, error)
	VerifyRefreshToken(string) (bool, error)
}
