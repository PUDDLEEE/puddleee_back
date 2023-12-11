package jwtAuth

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJWTAuthService_CreateAccessToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Creating Token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				ok, err := service.VerifyAccessToken(*token)

				require.NoError(t, err)
				require.Equal(t, true, ok)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}

func TestJWTAuthService_CreateRefreshToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Creating Refresh Token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateRefreshToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				ok, err := service.VerifyRefreshToken(*token)

				require.NoError(t, err)
				require.Equal(t, true, ok)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}

func TestJWTAuthService_ParseAccessToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Parse Access Token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsedToken, err := service.ParseAccessToken(*token)
				translatedToken, _ := parsedToken.Claims.(*AuthTokenClaims)
				require.NoError(t, err)
				require.Equal(t, true, parsedToken.Valid)
				require.Equal(t, 1, translatedToken.ID)
			},
		},
		{
			name: "Parse Invalid Access Token",
			expect: func(t *testing.T, service *JwtAuthService) {

				parsedToken, err := service.ParseAccessToken("asdsadlkasd")

				require.Error(t, err)
				require.Equal(t, (*jwt.Token)(nil), parsedToken)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}
func TestJWTAuthService_ParseRefreshToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Parse Access Token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateRefreshToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsedToken, err := service.ParseRefreshToken(*token)
				translatedToken, _ := parsedToken.Claims.(*AuthTokenClaims)
				require.NoError(t, err)
				require.Equal(t, true, parsedToken.Valid)
				require.Equal(t, 1, translatedToken.ID)
			},
		},
		{
			name: "Parse Invalid Access Token",
			expect: func(t *testing.T, service *JwtAuthService) {

				parsedToken, err := service.ParseRefreshToken("asdsadlkasd")

				require.Error(t, err)
				require.Equal(t, (*jwt.Token)(nil), parsedToken)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}

func TestJWTAuthService_VerifyAccessToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Verify valid access token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsedToken, err := service.VerifyAccessToken(*token)

				require.NoError(t, err)
				require.Equal(t, true, parsedToken)
			},
		},
		{
			name: "Verify invalid access token",
			expect: func(t *testing.T, service *JwtAuthService) {

				parsedToken, err := service.VerifyAccessToken("asdkalsdkjasld")

				require.Error(t, err)
				require.Equal(t, false, parsedToken)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}

func TestJWTAuthService_VerifyRefreshToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Verify valid refresh token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				token, err := service.CreateRefreshToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsedToken, err := service.VerifyRefreshToken(*token)

				require.NoError(t, err)
				require.Equal(t, true, parsedToken)
			},
		},
		{
			name: "Verify invalid refresh token",
			expect: func(t *testing.T, service *JwtAuthService) {

				parsedToken, err := service.VerifyRefreshToken("asdkalsdkjasld")

				require.Error(t, err)
				require.Equal(t, false, parsedToken)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}

func TestJwtAuthService_ReassignAccessToken(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *JwtAuthService)
		expect func(*testing.T, *JwtAuthService)
	}{
		{
			name: "Reassign valid access token",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				accessToken, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, accessToken)

				refreshToken, err := service.CreateRefreshToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, refreshToken)

				newRefreshToken, err := service.ReassignAccessToken(*accessToken, *refreshToken)

				require.NoError(t, err)
				require.NotEqual(t, refreshToken, newRefreshToken)
			},
		},
		{
			name: "Reassign with incorrect refreshToken",
			expect: func(t *testing.T, service *JwtAuthService) {
				claims := AuthTokenClaims{
					ID: 1,
				}
				otherClaims := AuthTokenClaims{
					ID: 2,
				}
				accessToken, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, accessToken)

				refreshToken, err := service.CreateRefreshToken(otherClaims)
				require.NoError(t, err)
				require.NotEmpty(t, refreshToken)

				_, err = service.ReassignAccessToken(*accessToken, *refreshToken)

				require.Error(t, err)
			},
		},
		{
			name: "Reassign with invalid accessToken",
			expect: func(t *testing.T, service *JwtAuthService) {

				otherClaims := AuthTokenClaims{
					ID: 2,
				}

				refreshToken, err := service.CreateRefreshToken(otherClaims)
				require.NoError(t, err)
				require.NotEmpty(t, refreshToken)

				_, err = service.ReassignAccessToken("asdklasdjs", *refreshToken)

				require.Error(t, err)
			},
		},
		{
			name: "Reassign with invalid refreshToken",
			expect: func(t *testing.T, service *JwtAuthService) {

				claims := AuthTokenClaims{
					ID: 2,
				}

				accessToken, err := service.CreateAccessToken(claims)
				require.NoError(t, err)
				require.NotEmpty(t, accessToken)

				_, err = service.ReassignAccessToken(*accessToken, "asldsajksdl")

				require.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 5
			b := make([]byte, n)
			if _, err := rand.Read(b); err != nil {
				t.Fatal(err)
			}
			key := fmt.Sprintf("%x", b)
			service := NewJwtAuthService(key)
			tt.expect(t, service)
		})
	}
}
