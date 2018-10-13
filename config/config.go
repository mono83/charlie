package config

// Config is main configuration structure
type Config struct {
	Auth Auth `ini:"auth"`
}

// Auth represents `auth` section of configuration
type Auth struct {
	Github string `ini:"github"`
}
