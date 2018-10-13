package ini

import (
	"github.com/mono83/charlie/config"
	"gopkg.in/ini.v1"
)

// GetDefaultConfig returns configuration structure read from file named `config.ini`
func GetDefaultConfig() (config.Config, error) {
	return loadConfig("config.ini")
}

// GetConfigFromSource returns configuration structure read from specified ini file
func GetConfigFromSource(file string) (config.Config, error) {
	return loadConfig(file)
}

// loadConfig loads config from specified file
func loadConfig(file string) (config.Config, error) {
	cfg, err := ini.Load(file)
	if err != nil {
		return config.Config{}, err
	}
	c := new(config.Config)
	err = cfg.MapTo(c)
	return *c, err
}
