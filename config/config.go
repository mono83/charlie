package config

type Config struct {
	Auth auth
}

type auth struct {
	Github string
}
