package config

type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type Jwt struct {
	SecretKey string
}

type Mail struct {
	Password string
	From     string
	Server   string
	Port     string
}

type Config struct {
	App    App
	Server Server
	Jwt    Jwt
	Mail   Mail
}
