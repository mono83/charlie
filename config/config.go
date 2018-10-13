package config

type Config struct {
	Auth Auth `ini:"auth"`
}

type Auth struct {
	Github string `ini:"github"`
}

// ConfigProvider is used to provide configuration values from different sources (e.g. .ini files, DB, etc.).
type ConfigProvider interface {
	// GetConfig returns configuration structure
	GetConfig() (Config, error)
}
