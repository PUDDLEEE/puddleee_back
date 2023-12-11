package auth

import "github.com/golang-jwt/jwt/v5"

type AuthTokenClaims struct {
	ID string `json:"id"`
	jwt.Claims
}
