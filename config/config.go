package config

type Config struct {
	Auth Auth `ini:"auth"`
}

type Auth struct {
	Github string `ini:"github"`
}
