package authdto

type SigninDTO struct {
	Email    string
	Password string
}

type SigninOutputDTO struct {
	AccessToken  string
	RefreshToken string
}
