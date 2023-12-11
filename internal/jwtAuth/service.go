package auth

import "github.com/golang-jwt/jwt/v5"

type AuthService struct {
	secretKey string
}

func (a *AuthService) CreateAccessToken(claims AuthTokenClaims) (*string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := accessToken.SignedString([]byte(a.secretKey))
	if err != nil {
		return nil, err
	}

	return &token, err
}

func (a *AuthService) CreateRefreshToken(claims jwt.Claims) (*string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := refreshToken.SignedString([]byte(a.secretKey))

	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (a *AuthService) ReassignAccessToken(refreshToken string) (*string, error) {
	return nil, nil
}

func (a *AuthService) VerifyAccessToken(accessToken string) (bool, error) {
	return true, nil
}

func (a *AuthService) VerifyRefreshToken(refreshToken string) (bool, error) {
	return true, nil
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{secretKey: secretKey}
}
