package config

type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type Config struct {
	App    App
	Server Server
}
