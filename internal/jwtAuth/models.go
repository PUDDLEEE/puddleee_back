package jwtAuth

import "github.com/golang-jwt/jwt/v5"

type AuthTokenClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
