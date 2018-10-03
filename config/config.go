package config

// ConfigProvider is used to provide configuration values from different sources (e.g. .ini files, DB, etc.).
type ConfigProvider interface {

	// GetConfig return configuration value (string) for specified input key
	GetConfig(key string) (string, error)
}
